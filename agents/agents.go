// DO NOT EDIT!
// Code generated by mage agents command
// source caniuse/data.json

// Package agents exposes details about all common web browsers. This uses data
// from caniuse-db.
package agents

// Agent contains details about a web browser.
type Agent struct {
	Name                string
	Browser             string
	Abbr                string
	Prefix              string
	Type                string
	UsageGlobal         map[string]float64
	Versions            []string
	DataPrefixEceptions map[string]string
}

// AgentsMap is mapping of browser name to browser details.
var AgentsMap = map[string]Agent{
	"and_chr": ChromeForAndroid,
	"and_ff":  FirefoxForAndroid,
	"and_qq":  QQForAndroid,
	"and_uc":  UCForAndroid,
	"android": Android,
	"baidu":   Baidu,
	"bb":      BlackBerry,
	"chrome":  Chrome,
	"edge":    Edge,
	"firefox": Firefox,
	"ie":      InternetExplorer,
	"ie_mob":  InternetExplorerMobile,
	"ios_saf": IOSSafari,
	"op_mini": OperaMini,
	"op_mob":  OperaMobile,
	"opera":   Opera,
	"safari":  Safari,
	"samsung": Samsung,
}
var (

	// ChromeForAndroid is Chrome for Android browser
	ChromeForAndroid = Agent{
		Name:        "and_chr",
		Browser:     "Chrome for Android",
		Abbr:        "Chr/And.",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"64": 30.4129},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "64", "", "", ""},
	}

	// FirefoxForAndroid is Firefox for Android browser
	FirefoxForAndroid = Agent{
		Name:        "and_ff",
		Browser:     "Firefox for Android",
		Abbr:        "FF/And.",
		Prefix:      "moz",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"57": 0.25047},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "57", "", "", ""},
	}

	// QQForAndroid is QQ Browser browser
	QQForAndroid = Agent{
		Name:        "and_qq",
		Browser:     "QQ Browser",
		Abbr:        "QQ",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"1.2": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "1.2", "", "", ""},
	}

	// UCForAndroid is UC Browser for Android browser
	UCForAndroid = Agent{
		Name:                "and_uc",
		Browser:             "UC Browser for Android",
		Abbr:                "UC",
		Prefix:              "webkit",
		Type:                "mobile",
		UsageGlobal:         map[string]float64{"11.8": 7.50853},
		Versions:            []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "11.8", "", "", ""},
		DataPrefixEceptions: map[string]string{"11.8": "webkit"}}

	// Android is Android Browser browser
	Android = Agent{
		Name:        "android",
		Browser:     "Android Browser",
		Abbr:        "And.",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"2.1": 0, "2.2": 0, "2.3": 0, "3": 0, "4": 0, "4.1": 0.067427, "4.2-4.3": 0.205212, "4.4": 0.578992, "4.4.3-4.4.4": 0.35619, "62": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "2.1", "2.2", "2.3", "3", "4", "4.1", "4.2-4.3", "4.4", "4.4.3-4.4.4", "62", "", "", ""},
	}

	// Baidu is Baidu Browser browser
	Baidu = Agent{
		Name:        "baidu",
		Browser:     "Baidu Browser",
		Abbr:        "baidu",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"7.12": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "7.12", "", "", ""},
	}

	// BlackBerry is Blackberry Browser browser
	BlackBerry = Agent{
		Name:        "bb",
		Browser:     "Blackberry Browser",
		Abbr:        "BB",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"10": 0.0534336, "7": 0.0133584},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "7", "10", "", "", ""},
	}

	// Chrome is Chrome browser
	Chrome = Agent{
		Name:        "chrome",
		Browser:     "Chrome",
		Abbr:        "Chr.",
		Prefix:      "webkit",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"10": 0.004534, "11": 0.008866, "12": 0.004283, "13": 0.004879, "14": 0.004706, "15": 0.009154, "16": 0.004393, "17": 0.004393, "18": 0.017732, "19": 0.004418, "20": 0.004393, "21": 0.004433, "22": 0.017732, "23": 0.008786, "24": 0.013299, "25": 0.008866, "26": 0.008866, "27": 0.008866, "28": 0.008866, "29": 0.363506, "30": 0.017732, "31": 0.026598, "32": 0.008866, "33": 0.013299, "34": 0.022165, "35": 0.013299, "36": 0.017732, "37": 0.017732, "38": 0.039897, "39": 0.017732, "4": 0.004706, "40": 0.017732, "41": 0.022165, "42": 0.022165, "43": 0.062062, "44": 0.013299, "45": 0.035464, "46": 0.022165, "47": 0.04433, "48": 0.124124, "49": 0.820105, "5": 0.004879, "50": 0.026598, "51": 0.053196, "52": 0.053196, "53": 0.026598, "54": 0.084227, "55": 0.359073, "56": 0.115258, "57": 0.097526, "58": 0.164021, "59": 0.115258, "6": 0.004879, "60": 0.159588, "61": 0.172887, "62": 0.203918, "63": 1.29887, "64": 15.2229, "65": 8.28528, "66": 0.053196, "67": 0.022165, "68": 0, "7": 0.005591, "8": 0.005591, "9": 0.005591},
		Versions:    []string{"4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68"},
	}

	// Edge is Edge browser
	Edge = Agent{
		Name:        "edge",
		Browser:     "Edge",
		Abbr:        "Edge",
		Prefix:      "ms",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"12": 0.026598, "13": 0.035464, "14": 0.097526, "15": 0.128557, "16": 1.57371, "17": 0.022165, "18": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "12", "13", "14", "15", "16", "17", "18", ""},
	}

	// Firefox is Firefox browser
	Firefox = Agent{
		Name:        "firefox",
		Browser:     "Firefox",
		Abbr:        "FF",
		Prefix:      "moz",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"10": 0.004283, "11": 0.004433, "12": 0.004418, "13": 0.004486, "14": 0.00453, "15": 0.004433, "16": 0.004433, "17": 0.004349, "18": 0.004393, "19": 0.004443, "2": 0.004433, "20": 0.004283, "21": 0.004418, "22": 0.004393, "23": 0.004433, "24": 0.008786, "25": 0.008836, "26": 0.004393, "27": 0.004393, "28": 0.004418, "29": 0.008866, "3": 0.004433, "3.5": 0.008786, "3.6": 0.008866, "30": 0.004433, "31": 0.008866, "32": 0.004433, "33": 0.004433, "34": 0.008866, "35": 0.008866, "36": 0.013299, "37": 0.008866, "38": 0.039897, "39": 0.008866, "4": 0.013299, "40": 0.013299, "41": 0.013299, "42": 0.008866, "43": 0.031031, "44": 0.075361, "45": 0.035464, "46": 0.008866, "47": 0.053196, "48": 0.137423, "49": 0.022165, "5": 0.004879, "50": 0.035464, "51": 0.062062, "52": 0.4433, "53": 0.026598, "54": 0.035464, "55": 0.04433, "56": 0.119691, "57": 0.106392, "58": 1.98155, "59": 1.57371, "6": 0.020136, "60": 0.048763, "61": 0, "7": 0.005725, "8": 0.008866, "9": 0.00533},
		Versions:    []string{"", "", "2", "3", "3.5", "3.6", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", ""},
	}

	// InternetExplorer is IE browser
	InternetExplorer = Agent{
		Name:        "ie",
		Browser:     "IE",
		Abbr:        "IE",
		Prefix:      "ms",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"10": 0.113991, "11": 2.76315, "5.5": 0.009298, "6": 0.00911931, "7": 0.013679, "8": 0.173267, "9": 0.118551},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "5.5", "6", "7", "8", "9", "10", "11", "", "", ""},
	}

	// InternetExplorerMobile is IE Mobile browser
	InternetExplorerMobile = Agent{
		Name:        "ie_mob",
		Browser:     "IE Mobile",
		Abbr:        "IE.Mob",
		Prefix:      "ms",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"10": 0.033396, "11": 0.178112},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "10", "11", "", "", ""},
	}

	// IOSSafari is iOS Safari browser
	IOSSafari = Agent{
		Name:        "ios_saf",
		Browser:     "iOS Safari",
		Abbr:        "iOS",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"10.0-10.2": 0.428659, "10.3": 0.993418, "11.0-11.2": 8.61391, "11.3": 0, "3.2": 0.00107165, "4.0-4.1": 0, "4.2-4.3": 0.00107165, "5.0-5.1": 0.0192897, "6.0-6.1": 0.0150031, "7.0-7.1": 0.0525108, "8": 0.0150031, "8.1-8.4": 0.079302, "9.0-9.2": 0.0600123, "9.3": 0.429731},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "3.2", "4.0-4.1", "4.2-4.3", "5.0-5.1", "6.0-6.1", "7.0-7.1", "8", "8.1-8.4", "9.0-9.2", "9.3", "10.0-10.2", "10.3", "11.0-11.2", "11.3", "", "", ""},
	}

	// OperaMini is Opera Mini browser
	OperaMini = Agent{
		Name:        "op_mini",
		Browser:     "Opera Mini",
		Abbr:        "O.Mini",
		Prefix:      "o",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"all": 2.66111},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "all", "", "", ""},
	}

	// OperaMobile is Opera Mobile browser
	OperaMobile = Agent{
		Name:                "op_mob",
		Browser:             "Opera Mobile",
		Abbr:                "O.Mob",
		Prefix:              "o",
		Type:                "mobile",
		UsageGlobal:         map[string]float64{"10": 0, "11": 0, "11.1": 0, "11.5": 0, "12": 0, "12.1": 0, "37": 0.0111391},
		Versions:            []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "10", "11", "11.1", "11.5", "12", "12.1", "37", "", "", ""},
		DataPrefixEceptions: map[string]string{"37": "webkit"}}

	// Opera is Opera browser
	Opera = Agent{
		Name:                "opera",
		Browser:             "Opera",
		Abbr:                "Op.",
		Prefix:              "webkit",
		Type:                "desktop",
		UsageGlobal:         map[string]float64{"10.0-10.1": 0, "10.5": 0.008392, "10.6": 0.004706, "11": 0.016581, "11.1": 0.006229, "11.5": 0.004879, "11.6": 0.008786, "12": 0.008786, "12.1": 0.04433, "15": 0.00685, "16": 0.00685, "17": 0.00685, "18": 0.005014, "19": 0.006015, "20": 0.004879, "21": 0.006597, "22": 0.006597, "23": 0.013434, "24": 0.006702, "25": 0.006015, "26": 0.005595, "27": 0.004393, "28": 0.008698, "29": 0.004879, "30": 0.004879, "31": 0.004433, "32": 0.005152, "33": 0.005014, "34": 0.009758, "35": 0.004879, "36": 0.039897, "37": 0.004283, "38": 0.004367, "39": 0.004534, "40": 0.004367, "41": 0.004227, "42": 0.004418, "43": 0.008668, "44": 0.004227, "45": 0.008866, "46": 0.004433, "47": 0.004433, "48": 0.004433, "49": 0.008866, "50": 0.048763, "51": 0.740311, "52": 0.026598, "9": 0.0082, "9.5-9.6": 0.00685},
		Versions:            []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "9", "9.5-9.6", "10.0-10.1", "10.5", "10.6", "11", "11.1", "11.5", "11.6", "12", "12.1", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", ""},
		DataPrefixEceptions: map[string]string{"10.0-10.1": "o", "10.5": "o", "10.6": "o", "11": "o", "11.1": "o", "11.5": "o", "11.6": "o", "12": "o", "12.1": "o", "9": "o", "9.5-9.6": "o"}}

	// Safari is Safari browser
	Safari = Agent{
		Name:        "safari",
		Browser:     "Safari",
		Abbr:        "Saf.",
		Prefix:      "webkit",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"10": 0.097526, "10.1": 0.319176, "11": 1.56928, "11.1": 0.022165, "3.1": 0, "3.2": 0.008692, "4": 0, "5": 0.013299, "5.1": 0.070928, "6": 0.004349, "6.1": 0.013299, "7": 0.008866, "7.1": 0.004283, "8": 0.031031, "9": 0.031031, "9.1": 0.164021, "TP": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "3.1", "3.2", "4", "5", "5.1", "6", "6.1", "7", "7.1", "8", "9", "9.1", "10", "10.1", "11", "11.1", "TP", "", ""},
	}

	// Samsung is Samsung Internet browser
	Samsung = Agent{
		Name:        "samsung",
		Browser:     "Samsung Internet",
		Abbr:        "SS",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"4": 0.943211, "5": 0.228029, "6.2": 1.71022},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "4", "5", "6.2", "", "", ""},
	}
)

// All is a helper function that returns a list of all agents.
func All() []Agent {
	return []Agent{ChromeForAndroid, FirefoxForAndroid, QQForAndroid, UCForAndroid, Android, Baidu, BlackBerry, Chrome, Edge, Firefox, InternetExplorer, InternetExplorerMobile, IOSSafari, OperaMini, OperaMobile, Opera, Safari, Samsung}
}
