package searchselect

// todo: move this to a file that's more common when i figure out where it should belong
func GetCountryOptions() []Option {
	return []Option{
		{Value: "af", Text: "Afghanistan"},
		{Value: "al", Text: "Albania"},
		{Value: "dz", Text: "Algeria"},
		{Value: "ad", Text: "Andorra"},
		{Value: "ao", Text: "Angola"},
		{Value: "ag", Text: "Antigua and Barbuda"},
		{Value: "ar", Text: "Argentina"},
		{Value: "am", Text: "Armenia"},
		{Value: "au", Text: "Australia"},
		{Value: "at", Text: "Austria"},
		{Value: "az", Text: "Azerbaijan"},
		{Value: "bs", Text: "Bahamas"},
		{Value: "bh", Text: "Bahrain"},
		{Value: "bd", Text: "Bangladesh"},
		{Value: "bb", Text: "Barbados"},
		{Value: "by", Text: "Belarus"},
		{Value: "be", Text: "Belgium"},
		{Value: "bz", Text: "Belize"},
		{Value: "bj", Text: "Benin"},
		{Value: "bt", Text: "Bhutan"},
		{Value: "bo", Text: "Bolivia"},
		{Value: "ba", Text: "Bosnia and Herzegovina"},
		{Value: "bw", Text: "Botswana"},
		{Value: "br", Text: "Brazil"},
		{Value: "bn", Text: "Brunei"},
		{Value: "bg", Text: "Bulgaria"},
		{Value: "bf", Text: "Burkina Faso"},
		{Value: "bi", Text: "Burundi"},
		{Value: "kh", Text: "Cambodia"},
		{Value: "cm", Text: "Cameroon"},
		{Value: "ca", Text: "Canada"},
		{Value: "cv", Text: "Cape Verde"},
		{Value: "cf", Text: "Central African Republic"},
		{Value: "td", Text: "Chad"},
		{Value: "cl", Text: "Chile"},
		{Value: "cn", Text: "China"},
		{Value: "co", Text: "Colombia"},
		{Value: "km", Text: "Comoros"},
		{Value: "cg", Text: "Congo"},
		{Value: "cd", Text: "Congo, Democratic Republic"},
		{Value: "cr", Text: "Costa Rica"},
		{Value: "ci", Text: "Côte d'Ivoire"},
		{Value: "hr", Text: "Croatia"},
		{Value: "cu", Text: "Cuba"},
		{Value: "cy", Text: "Cyprus"},
		{Value: "cz", Text: "Czech Republic"},
		{Value: "dk", Text: "Denmark"},
		{Value: "dj", Text: "Djibouti"},
		{Value: "dm", Text: "Dominica"},
		{Value: "do", Text: "Dominican Republic"},
		{Value: "ec", Text: "Ecuador"},
		{Value: "eg", Text: "Egypt"},
		{Value: "sv", Text: "El Salvador"},
		{Value: "gq", Text: "Equatorial Guinea"},
		{Value: "er", Text: "Eritrea"},
		{Value: "ee", Text: "Estonia"},
		{Value: "et", Text: "Ethiopia"},
		{Value: "fj", Text: "Fiji"},
		{Value: "fi", Text: "Finland"},
		{Value: "fr", Text: "France"},
		{Value: "ga", Text: "Gabon"},
		{Value: "gm", Text: "Gambia"},
		{Value: "ge", Text: "Georgia"},
		{Value: "de", Text: "Germany"},
		{Value: "gh", Text: "Ghana"},
		{Value: "gr", Text: "Greece"},
		{Value: "gd", Text: "Grenada"},
		{Value: "gt", Text: "Guatemala"},
		{Value: "gn", Text: "Guinea"},
		{Value: "gw", Text: "Guinea-Bissau"},
		{Value: "gy", Text: "Guyana"},
		{Value: "ht", Text: "Haiti"},
		{Value: "hn", Text: "Honduras"},
		{Value: "hu", Text: "Hungary"},
		{Value: "is", Text: "Iceland"},
		{Value: "in", Text: "India"},
		{Value: "id", Text: "Indonesia"},
		{Value: "ir", Text: "Iran"},
		{Value: "iq", Text: "Iraq"},
		{Value: "ie", Text: "Ireland"},
		{Value: "il", Text: "Israel"},
		{Value: "it", Text: "Italy"},
		{Value: "jm", Text: "Jamaica"},
		{Value: "jp", Text: "Japan"},
		{Value: "jo", Text: "Jordan"},
		{Value: "kz", Text: "Kazakhstan"},
		{Value: "ke", Text: "Kenya"},
		{Value: "ki", Text: "Kiribati"},
		{Value: "kp", Text: "Korea, North"},
		{Value: "kr", Text: "Korea, South"},
		{Value: "kw", Text: "Kuwait"},
		{Value: "kg", Text: "Kyrgyzstan"},
		{Value: "la", Text: "Laos"},
		{Value: "lv", Text: "Latvia"},
		{Value: "lb", Text: "Lebanon"},
		{Value: "ls", Text: "Lesotho"},
		{Value: "lr", Text: "Liberia"},
		{Value: "ly", Text: "Libya"},
		{Value: "li", Text: "Liechtenstein"},
		{Value: "lt", Text: "Lithuania"},
		{Value: "lu", Text: "Luxembourg"},
		{Value: "mk", Text: "Macedonia"},
		{Value: "mg", Text: "Madagascar"},
		{Value: "mw", Text: "Malawi"},
		{Value: "my", Text: "Malaysia"},
		{Value: "mv", Text: "Maldives"},
		{Value: "ml", Text: "Mali"},
		{Value: "mt", Text: "Malta"},
		{Value: "mh", Text: "Marshall Islands"},
		{Value: "mr", Text: "Mauritania"},
		{Value: "mu", Text: "Mauritius"},
		{Value: "mx", Text: "Mexico"},
		{Value: "fm", Text: "Micronesia"},
		{Value: "md", Text: "Moldova"},
		{Value: "mc", Text: "Monaco"},
		{Value: "mn", Text: "Mongolia"},
		{Value: "me", Text: "Montenegro"},
		{Value: "ma", Text: "Morocco"},
		{Value: "mz", Text: "Mozambique"},
		{Value: "mm", Text: "Myanmar"},
		{Value: "na", Text: "Namibia"},
		{Value: "nr", Text: "Nauru"},
		{Value: "np", Text: "Nepal"},
		{Value: "nl", Text: "Netherlands"},
		{Value: "nz", Text: "New Zealand"},
		{Value: "ni", Text: "Nicaragua"},
		{Value: "ne", Text: "Niger"},
		{Value: "ng", Text: "Nigeria"},
		{Value: "no", Text: "Norway"},
		{Value: "om", Text: "Oman"},
		{Value: "pk", Text: "Pakistan"},
		{Value: "pw", Text: "Palau"},
		{Value: "pa", Text: "Panama"},
		{Value: "pg", Text: "Papua New Guinea"},
		{Value: "py", Text: "Paraguay"},
		{Value: "pe", Text: "Peru"},
		{Value: "ph", Text: "Philippines"},
		{Value: "pl", Text: "Poland"},
		{Value: "pt", Text: "Portugal"},
		{Value: "qa", Text: "Qatar"},
		{Value: "ro", Text: "Romania"},
		{Value: "ru", Text: "Russia"},
		{Value: "rw", Text: "Rwanda"},
		{Value: "kn", Text: "Saint Kitts and Nevis"},
		{Value: "lc", Text: "Saint Lucia"},
		{Value: "vc", Text: "Saint Vincent and the Grenadines"},
		{Value: "ws", Text: "Samoa"},
		{Value: "sm", Text: "San Marino"},
		{Value: "st", Text: "São Tomé and Príncipe"},
		{Value: "sa", Text: "Saudi Arabia"},
		{Value: "sn", Text: "Senegal"},
		{Value: "rs", Text: "Serbia"},
		{Value: "sc", Text: "Seychelles"},
		{Value: "sl", Text: "Sierra Leone"},
		{Value: "sg", Text: "Singapore"},
		{Value: "sk", Text: "Slovakia"},
		{Value: "si", Text: "Slovenia"},
		{Value: "sb", Text: "Solomon Islands"},
		{Value: "so", Text: "Somalia"},
		{Value: "za", Text: "South Africa"},
		{Value: "ss", Text: "South Sudan"},
		{Value: "es", Text: "Spain"},
		{Value: "lk", Text: "Sri Lanka"},
		{Value: "sd", Text: "Sudan"},
		{Value: "sr", Text: "Suriname"},
		{Value: "sz", Text: "Swaziland"},
		{Value: "se", Text: "Sweden"},
		{Value: "ch", Text: "Switzerland"},
		{Value: "sy", Text: "Syria"},
		{Value: "tw", Text: "Taiwan"},
		{Value: "tj", Text: "Tajikistan"},
		{Value: "tz", Text: "Tanzania"},
		{Value: "th", Text: "Thailand"},
		{Value: "tl", Text: "Timor-Leste"},
		{Value: "tg", Text: "Togo"},
		{Value: "to", Text: "Tonga"},
		{Value: "tt", Text: "Trinidad and Tobago"},
		{Value: "tn", Text: "Tunisia"},
		{Value: "tr", Text: "Turkey"},
		{Value: "tm", Text: "Turkmenistan"},
		{Value: "tv", Text: "Tuvalu"},
		{Value: "ug", Text: "Uganda"},
		{Value: "ua", Text: "Ukraine"},
		{Value: "ae", Text: "United Arab Emirates"},
		{Value: "gb", Text: "United Kingdom"},
		{Value: "us", Text: "United States"},
		{Value: "uy", Text: "Uruguay"},
		{Value: "uz", Text: "Uzbekistan"},
		{Value: "vu", Text: "Vanuatu"},
		{Value: "va", Text: "Vatican City"},
		{Value: "ve", Text: "Venezuela"},
		{Value: "vn", Text: "Vietnam"},
		{Value: "ye", Text: "Yemen"},
		{Value: "zm", Text: "Zambia"},
		{Value: "zw", Text: "Zimbabwe"},
	}
}
