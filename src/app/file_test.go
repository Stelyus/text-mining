package app

import ("testing")

func TestDeserialize(test *testing.T) {
	test.Log("Launching test on Deserialize")
	//trie := Deserialize("../../ressources/dict_words.bin")
	// test.Errorf("Sum was incorrect, got: %d, want: %d.", 1, 10)

}

func TestGetDistance(test *testing.T) {
	test.Log("Launching test on GetDistance")
	trie := Deserialize("../../ressources/dict_words.bin")

	GetDistance(trie, "nabilo", 1)

	// echo "approx nabilo 1"
	ref := "[{\"word\":\"nabilo\",\"freq\":970,\"distance\":0},{\"word\":\"nabil\",\"freq\":365545,\"distance\":1},{\"word\":\"nabila\",\"freq\":102158,\"distance\":1},{\"word\":\"nailo\",\"freq\":19881,\"distance\":1},{\"word\":\"nobilo\",\"freq\":8739,\"distance\":1},{\"word\":\"nabile\",\"freq\":7070,\"distance\":1},{\"word\":\"nabilou\",\"freq\":4130,\"distance\":1},{\"word\":\"nabiloo\",\"freq\":2057,\"distance\":1},{\"word\":\"nabili\",\"freq\":1862,\"distance\":1},{\"word\":\"babilo\",\"freq\":1461,\"distance\":1},{\"word\":\"nabilos\",\"freq\":1448,\"distance\":1},{\"word\":\"nabilon\",\"freq\":873,\"distance\":1},{\"word\":\"nabil2\",\"freq\":579,\"distance\":1},{\"word\":\"nabio\",\"freq\":429,\"distance\":1},{\"word\":\"abilo\",\"freq\":337,\"distance\":1},{\"word\":\"nadilo\",\"freq\":310,\"distance\":1},{\"word\":\"nablo\",\"freq\":289,\"distance\":1},{\"word\":\"nabill\",\"freq\":284,\"distance\":1},{\"word\":\"fabilo\",\"freq\":261,\"distance\":1},{\"word\":\"kabilo\",\"freq\":244,\"distance\":1}]"

	// echo "approx 0 test"
	ref1 := "[{\"word\":\"test\",\"freq\":49216987,\"distance\":0}]"

	// echo "approx 1 utard"
	ref2 := "[{\"word\":\"utard\",\"freq\":5044,\"distance\":0},{\"word\":\"tard\",\"freq\":22348107,\"distance\":1},{\"word\":\"itard\",\"freq\":8859,\"distance\":1},{\"word\":\"stard\",\"freq\":6646,\"distance\":1},{\"word\":\"etard\",\"freq\":6045,\"distance\":1},{\"word\":\"utgard\",\"freq\":5709,\"distance\":1},{\"word\":\"autard\",\"freq\":5405,\"distance\":1},{\"word\":\"jutard\",\"freq\":5217,\"distance\":1},{\"word\":\"tutard\",\"freq\":4876,\"distance\":1},{\"word\":\"dutard\",\"freq\":3627,\"distance\":1},{\"word\":\"otard\",\"freq\":3426,\"distance\":1},{\"word\":\"utara\",\"freq\":3030,\"distance\":1},{\"word\":\"butard\",\"freq\":2557,\"distance\":1},{\"word\":\"rtard\",\"freq\":1389,\"distance\":1},{\"word\":\"utar\",\"freq\":1302,\"distance\":1},{\"word\":\"lutard\",\"freq\":1192,\"distance\":1},{\"word\":\"ptard\",\"freq\":1170,\"distance\":1},{\"word\":\"rutard\",\"freq\":891,\"distance\":1},{\"word\":\"atard\",\"freq\":879,\"distance\":1},{\"word\":\"utari\",\"freq\":720,\"distance\":1},{\"word\":\"uard\",\"freq\":619,\"distance\":1},{\"word\":\"utad\",\"freq\":526,\"distance\":1},{\"word\":\"ttard\",\"freq\":415,\"distance\":1},{\"word\":\"cutard\",\"freq\":289,\"distance\":1},{\"word\":\"outard\",\"freq\":275,\"distance\":1},{\"word\":\"utare\",\"freq\":263,\"distance\":1},{\"word\":\"utarid\",\"freq\":244,\"distance\":1}]"


}