package cards

type Card struct {
	name         string // 7, 8, 9, 10, knight, jack , queen, king
	type_of_card string // clover, diamond, heart, spade, tarock
	img          string // "./img/heart5.png", ...
	power        int    // 1...8 for non tarock cards, 1...22 for tarocks
	points       int    // value of points at the end of the game
}

var clover1 = Card{"7", "clover", "./img/clover1.png", 1, 1}
var clover2 = Card{"8", "clover", "./img/clover2.png", 2, 1}
var clover3 = Card{"9", "clover", "./img/clover3.png", 3, 1}
var clover4 = Card{"10", "clover", "./img/clover4.png", 4, 1}
var clover5 = Card{"knight", "clover", "./img/clover5.png", 5, 2}
var clover6 = Card{"jack", "clover", "./img/clover6.png", 6, 3}
var clover7 = Card{"queen", "clover", "./img/clover7.png", 7, 4}
var clover8 = Card{"king", "clover", "./img/clover8.png", 8, 5}

var diamond1 = Card{"7", "diamond", "./img/diamond1.png", 1, 1}
var diamond2 = Card{"8", "diamond", "./img/diamond2.png", 2, 1}
var diamond3 = Card{"9", "diamond", "./img/diamond3.png", 3, 1}
var diamond4 = Card{"10", "diamond", "./img/diamond4.png", 4, 1}
var diamond5 = Card{"knight", "diamond", "./img/diamond5.png", 5, 2}
var diamond6 = Card{"jack", "diamond", "./img/diamond6.png", 6, 3}
var diamond7 = Card{"queen", "diamond", "./img/diamond7.png", 7, 4}
var diamond8 = Card{"king", "diamond", "./img/diamond8.png", 8, 5}

var heart1 = Card{"7", "heart", "./img/heart1.png", 1, 1}
var heart2 = Card{"8", "heart", "./img/heart2.png", 2, 1}
var heart3 = Card{"9", "heart", "./img/heart3.png", 3, 1}
var heart4 = Card{"10", "heart", "./img/heart4.png", 4, 1}
var heart5 = Card{"knight", "heart", "./img/heart5.png", 5, 2}
var heart6 = Card{"jack", "heart", "./img/heart6.png", 6, 3}
var heart7 = Card{"queen", "heart", "./img/heart7.png", 7, 4}
var heart8 = Card{"king", "heart", "./img/heart8.png", 8, 5}

var spade1 = Card{"7", "spade", "./img/spade1.png", 1, 1}
var spade2 = Card{"8", "spade", "./img/spade2.png", 2, 1}
var spade3 = Card{"9", "spade", "./img/spade3.png", 3, 1}
var spade4 = Card{"10", "spade", "./img/spade4.png", 4, 1}
var spade5 = Card{"knight", "spade", "./img/spade5.png", 5, 2}
var spade6 = Card{"jack", "spade", "./img/spade6.png", 6, 3}
var spade7 = Card{"queen", "spade", "./img/spade7.png", 7, 4}
var spade8 = Card{"king", "spade", "./img/spade8.png", 8, 5}

var tarock1 = Card{"mond", "tarock", "./img/tarock1.png", 1, 5}
var tarock2 = Card{"trump", "tarock", "./img/tarock2.png", 2, 1}
var tarock3 = Card{"trump", "tarock", "./img/tarock3.png", 3, 1}
var tarock4 = Card{"trump", "tarock", "./img/tarock4.png", 4, 1}
var tarock5 = Card{"trump", "tarock", "./img/tarock5.png", 5, 1}
var tarock6 = Card{"trump", "tarock", "./img/tarock6.png", 6, 1}
var tarock7 = Card{"trump", "tarock", "./img/tarock7.png", 7, 1}
var tarock8 = Card{"trump", "tarock", "./img/tarock8.png", 8, 1}
var tarock9 = Card{"trump", "tarock", "./img/tarock9.png", 9, 1}

var tarock10 = Card{"trump", "tarock", "./img/tarock10.png", 10, 1}
var tarock11 = Card{"trump", "tarock", "./img/tarock11.png", 11, 1}
var tarock12 = Card{"trump", "tarock", "./img/tarock12.png", 12, 1}
var tarock13 = Card{"trump", "tarock", "./img/tarock13.png", 13, 1}
var tarock14 = Card{"trump", "tarock", "./img/tarock14.png", 14, 1}

var tarock15 = Card{"trump", "tarock", "./img/tarock15.png", 15, 1}
var tarock16 = Card{"trump", "tarock", "./img/tarock16.png", 16, 1}
var tarock17 = Card{"trump", "tarock", "./img/tarock17.png", 17, 1}
var tarock18 = Card{"trump", "tarock", "./img/tarock18.png", 18, 1}
var tarock19 = Card{"trump", "tarock", "./img/tarock10.png", 19, 1}
var tarock20 = Card{"trump", "tarock", "./img/tarock20.png", 20, 1}
var tarock21 = Card{"mond", "tarock", "./img/tarock21.png", 21, 5}
var tarock22 = Card{"trump", "tarock", "./img/tarock22.png", 22, 5}

var Deck = []Card{
	clover1, clover2, clover3, clover4, clover5, clover6, clover7, clover8,
	diamond1, diamond2, diamond3, diamond4, diamond5, diamond6, diamond7, diamond8,
	heart1, heart2, heart3, heart4, heart5, heart6, heart7, heart8,
	spade1, spade2, spade3, spade4, spade5, spade6, spade7, spade8,
	tarock1, tarock2, tarock3, tarock4, tarock5, tarock6, tarock7, tarock8, tarock9, tarock10, tarock11, tarock12, tarock13, tarock14, tarock15, tarock16, tarock17, tarock18, tarock19, tarock20, tarock21, tarock22,
}
