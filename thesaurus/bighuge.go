package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type words struct {
	Syn []string `json:"syn"`
}
type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

// BigHuge ... APIキーを保持
type BigHuge struct {
	APIKey string
}

// Synonyms ... [Noun.Syn]と[Verb.Syn]のそれぞれを取得
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	// 単語（term）を受け取って、WebAPIを使って類語検索
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge: %qの類語検索に失敗しました： %v", term, err)
	}
	defer response.Body.Close()

	var data synonyms
	// HTTPレスポンスをJSONとしてパースし、synonyms型構造体のdataに格納
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns, nil
}
