// 中文单词学习程序
// author: baoqiang
// time: 2019-08-27 15:31
package app

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var d *Dictionary

func Server() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: ", os.Args[0], ":port\n")
	}

	port := os.Args[1]

	// dict path
	dictionaryPath := "data/cedict_ts.u8"
	d = new(Dictionary)

	// dict load
	d.Load(dictionaryPath)
	fmt.Println("Loaded dict ", len(d.Entries))

	// add handler
	http.HandleFunc("/", listFlashCards)
	http.HandleFunc("/wordlook", lookupWord)
	http.HandleFunc("/flashcards.html", listFlashCards)
	http.HandleFunc("/flashcardSets", manageFlashCards)
	http.HandleFunc("/searchWord", searchWord)
	http.HandleFunc("/addWord", addWord)
	http.HandleFunc("/newFlashCardSet", newFlashCardSet)

	// add file server
	fileServer := http.StripPrefix("/jscript/", http.FileServer(http.Dir("/data/jscript/")))
	http.Handle("/jscript/", fileServer)
	fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("/data/html/")))
	http.Handle("/html/", fileServer)

	err := http.ListenAndServe(port, nil)
	checkError(err)

}

// handle func
func lookupWord(rw http.ResponseWriter, req *http.Request) {
	word := req.FormValue("word")
	words := d.LookupEnglish(word)

	t := template.New("DictionaryEntry.html")
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})

	t, err := t.ParseFiles("html/DictionaryEntry.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(rw, words)

}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {
	// TODO
	//flashCardsNames := ListFlashCardsNames()
	flashCardsNames := ""

	t, err := template.ParseFiles("html/ListFlashcards.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(rw, flashCardsNames)
}

func manageFlashCards(rw http.ResponseWriter, req *http.Request) {
	set := req.FormValue("flashcardSets")
	order := req.FormValue("order")
	action := req.FormValue("submit")
	half := req.FormValue("half")

	fmt.Println("set chosen is", set)
	fmt.Println("order is", order)
	fmt.Println("action is", action)

	cardname := "flashcardSets/" + set
	fmt.Println("cardname", cardname, "action", action)

	if action == "Show card in set" {
		showFlashCards(rw, cardname, order, half)
	} else if action == "List words in set" {
		listWords(rw, cardname)
	} else if action == "Add cards to set" {
		addFlashCards(rw, set)
	}
}

func searchWord(rw http.ResponseWriter, req *http.Request) {
	word := req.FormValue("word")
	searchType := req.FormValue("searchType")
	cardName := req.FormValue("cardName")

	var words *Dictionary
	var dp []DictPlus

	if searchType == "english" {
		words = d.LookupEnglish(word)
		d1 := DictPlus{Dictionary: words, Word: word, CardName: cardName}
		dp = append(dp, d1)
	} else {
		words = d.LookupPinyin(word)
		numTrans := 0

		for _, entry := range words.Entries {
			numTrans += len(entry.Translations)
		}

		dp = make([]DictPlus, numTrans)
		idx := 0

		for _, entry := range words.Entries {
			for _, trans := range entry.Translations {
				dict := new(Dictionary)
				dict.Entries = make([]*Entry, 1)
				dp[idx] = DictPlus{
					Dictionary: dict,
					Word:       trans,
					CardName:   cardName,
				}

				idx++
			}
		}
	}

	// template
	t := template.New("ChooseDictionaryEntry.html")
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})

	t, err := t.ParseFiles("html/ChooseDictionaryEntry.html")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(rw, dp)

}

func addWord(rw http.ResponseWriter, req *http.Request) {
	url := req.URL
	fmt.Println("url", url.String())
	fmt.Println("query", url.RawQuery)

	word := req.FormValue("word")
	cardName := req.FormValue("cardname")
	simplified := req.FormValue("simplified")
	pinyin := req.FormValue("pinyin")
	traditional := req.FormValue("traditional")
	translations := req.FormValue("translations")

	fmt.Println("word is ", word, " card is ", cardName,
		" simplified is ", simplified, " pinyin is ", pinyin,
		" trad is ", traditional, " trans is ", translations)

	//TODO
	//AddFlashEntry(cardName, word, pinyin, simplified, traditional, translations)

	addFlashCards(rw, cardName)

}

func newFlashCardSet(rw http.ResponseWriter, req *http.Request) {
	defer http.Redirect(rw, req, "http://flashcards.html", 200)

	newSet := req.FormValue("NewFlashCard")
	fmt.Println("New cards", newSet)

	b, err := regexp.Match("[/$~]", []byte(newSet))
	if err != nil {
		return
	}

	if b {
		fmt.Println("no good string")
		return
	}

	// TODO
	//NewFlashCardSet(newSet)
	return
}

// others
func indexPage(rw http.ResponseWriter, req *http.Request) {
	index, _ := ioutil.ReadFile("html/index.html")
	rw.Write([]byte(index))
}

func addFlashCards(rw http.ResponseWriter, cardname string) {
	t, err := template.ParseFiles("html/AddWordToSet.html")
	if err != nil {
		fmt.Println("Parse err: ", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	//TODO
	//cards := GetFlashCardsByName(cardname,d)
	cards := ""
	t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute err: ", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func showFlashCards(rw http.ResponseWriter, cardname, order, half string) {
	fmt.Println("Loading card name", cardname)
	cards := new(FlashCards)
	//LoadJson(cardname,cards)

	if order == "Sequential" {
		cards.CardOrder = "SEQUENTIAL"
	} else {
		cards.CardOrder = "RANDOM"
	}

	fmt.Println("half is", half)
	if half == "Random" {
		cards.ShowHalf = "RANDOM_HALF"
	} else if half == "English" {
		cards.ShowHalf = "ENGLISH_HALF"
	} else {
		cards.ShowHalf = "CHINESE_HALF"
	}

	fmt.Println("loaded cards", len(cards.Cards))
	fmt.Println("Card name", cards.Name)

	t := template.New("ShowFlashcards.html")
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})
	t, err := t.ParseFiles("html/ShowFlashcards.html")

	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func listWords(rw http.ResponseWriter, cardname string) {
	fmt.Println("Loading card name", cardname)
	cards := new(FlashCards)

	// TODO
	//LoadJSON(cardname, cards)
	fmt.Println("loaded cards", len(cards.Cards))
	fmt.Println("Card name", cards.Name)

	t := template.New("ListWords.html")
	if t.Tree == nil {
		fmt.Println("New t is an incomplete or empty template")
	}
	t = t.Funcs(template.FuncMap{"pinyin": PinyinFormatter})
	t, err := t.ParseFiles("html/ListWords.html")
	if t.Tree == nil {
		fmt.Println("Parsed t is an incomplete or empty template")
	}

	if err != nil {
		fmt.Println("Parse error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(rw, cards)
	if err != nil {
		fmt.Println("Execute error " + err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("No error ")
}

type DictPlus struct {
	*Dictionary
	Word     string
	CardName string
}
