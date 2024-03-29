/*

  Sigma Virtual Machine
  Copyright (C) 2016  Dmitry Kolesnikov

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as published
  by the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package gen

import (
	"github.com/kshard/sigma/vm"
	"github.com/kshard/xsd"
)

func FactsIMDB(addr []vm.Addr) vm.Stream {
	return NewSubQ(addr, IMDB())
}

func IMDB() [][]xsd.Value {
	return [][]xsd.Value{
		// {"urn:person:100", "name", "James Cameron"},
		// {"urn:person:100", "born", "1954-08-16"},

		// {"urn:person:101", "name", "Arnold Schwarzenegger"},
		// {"urn:person:101", "born", "1947-07-30"},

		// {"urn:person:102", "name", "Linda Hamilton"},
		// {"urn:person:102", "born", "1956-09-26"},

		// {"urn:person:103", "name", "Michael Biehn"},
		// {"urn:person:103", "born", "1956-07-31"},

		// {"urn:person:104", "name", "Ted Kotcheff"},
		// {"urn:person:104", "born", "1931-04-07"},

		// {"urn:person:105", "name", "Sylvester Stallone"},
		// {"urn:person:105", "born", "1946-07-06"},

		// {"urn:person:106", "name", "Richard Crenna"},
		// {"urn:person:106", "born", "1926-11-30"},
		// {"urn:person:106", "death", "2003-01-17"},

		// {"urn:person:107", "name", "Brian Dennehy"},
		// {"urn:person:107", "born", "1938-07-09"},

		// {"urn:person:108", "name", "John McTiernan"},
		// {"urn:person:108", "born", "1951-01-08"},

		// {"urn:person:109", "name", "Elpidia Carrillo"},
		// {"urn:person:109", "born", "1961-08-16"},

		// {"urn:person:110", "name", "Carl Weathers"},
		// {"urn:person:110", "born", "1948-01-14"},

		// {"urn:person:111", "name", "Richard Donner"},
		// {"urn:person:111", "born", "1930-04-24"},

		// {"urn:person:112", "name", "Mel Gibson"},
		// {"urn:person:112", "born", "1956-01-03"},

		// {"urn:person:113", "name", "Danny Glover"},
		// {"urn:person:113", "born", "1946-07-22"},

		// {"urn:person:114", "name", "Gary Busey"},
		// {"urn:person:114", "born", "1944-07-29"},

		// {"urn:person:115", "name", "Paul Verhoeven"},
		// {"urn:person:115", "born", "1938-07-18"},

		// {"urn:person:116", "name", "Peter Weller"},
		// {"urn:person:116", "born", "1947-06-24"},

		// {"urn:person:117", "name", "Nancy Allen"},
		// {"urn:person:117", "born", "1950-06-24"},

		// {"urn:person:118", "name", "Ronny Cox"},
		// {"urn:person:118", "born", "1938-07-23"},

		// {"urn:person:119", "name", "Mark L. Lester"},
		// {"urn:person:119", "born", "1946-11-26"},

		// {"urn:person:120", "name", "Rae Dawn Chong"},
		// {"urn:person:120", "born", "1961-02-28"},

		// {"urn:person:121", "name", "Alyssa Milano"},
		// {"urn:person:121", "born", "1972-12-19"},

		// {"urn:person:122", "name", "Bruce Willis"},
		// {"urn:person:122", "born", "1955-03-19"},

		// {"urn:person:123", "name", "Alan Rickman"},
		// {"urn:person:123", "born", "1946-02-21"},

		// {"urn:person:124", "name", "Alexander Godunov"},
		// {"urn:person:124", "born", "1949-11-28"},
		// {"urn:person:124", "death", "1995-05-18"},

		// {"urn:person:125", "name", "Robert Patrick"},
		// {"urn:person:125", "born", "1958-11-05"},

		// {"urn:person:126", "name", "Edward Furlong"},
		// {"urn:person:126", "born", "1977-08-02"},

		// {"urn:person:127", "name", "Jonathan Mostow"},
		// {"urn:person:127", "born", "1961-11-28"},

		// {"urn:person:128", "name", "Nick Stahl"},
		// {"urn:person:128", "born", "1979-12-05"},

		// {"urn:person:129", "name", "Claire Danes"},
		// {"urn:person:129", "born", "1979-04-12"},

		// {"urn:person:130", "name", "George P. Cosmatos"},
		// {"urn:person:130", "born", "1941-01-04"},
		// {"urn:person:130", "death", "2005-04-19"},

		// {"urn:person:131", "name", "Charles Napier"},
		// {"urn:person:131", "born", "1936-04-12"},
		// {"urn:person:131", "death", "2011-10-05"},

		// {"urn:person:132", "name", "Peter MacDonald"},

		// {"urn:person:133", "name", "Marc de Jonge"},
		// {"urn:person:133", "born", "1949-02-16"},
		// {"urn:person:133", "death", "1996-06-06"},

		// {"urn:person:134", "name", "Stephen Hopkins"},

		// {"urn:person:135", "name", "Ruben Blades"},
		// {"urn:person:135", "born", "1948-07-16"},

		// {"urn:person:136", "name", "Joe Pesci"},
		// {"urn:person:136", "born", "1943-02-09"},

		// {"urn:person:137", "name", "Ridley Scott"},
		// {"urn:person:137", "born", "1937-11-30"},

		// {"urn:person:138", "name", "Tom Skerritt"},
		// {"urn:person:138", "born", "1933-08-25"},

		// {"urn:person:139", "name", "Sigourney Weaver"},
		// {"urn:person:139", "born", "1949-10-08"},

		// {"urn:person:140", "name", "Veronica Cartwright"},
		// {"urn:person:140", "born", "1949-04-20"},

		// {"urn:person:141", "name", "Carrie Henn"},

		// {"urn:person:142", "name", "George Miller"},
		// {"urn:person:142", "born", "1945-03-03"},

		// {"urn:person:143", "name", "Steve Bisley"},
		// {"urn:person:143", "born", "1951-12-26"},

		// {"urn:person:144", "name", "Joanne Samuel"},

		// {"urn:person:145", "name", "Michael Preston"},
		// {"urn:person:145", "born", "1938-05-14"},

		// {"urn:person:146", "name", "Bruce Spence"},
		// {"urn:person:146", "born", "1945-09-17"},

		// {"urn:person:147", "name", "George Ogilvie"},
		// {"urn:person:147", "born", "1931-03-05"},

		// {"urn:person:148", "name", "Tina Turner"},
		// {"urn:person:148", "born", "1939-11-26"},

		// {"urn:person:149", "name", "Sophie Marceau"},
		// {"urn:person:149", "born", "1966-11-17"},

		// {"urn:movie:200", "title", "The Terminator"},
		// {"urn:movie:200", "year", 1984},
		// {"urn:movie:200", "director", "urn:person:100"},
		// {"urn:movie:200", "cast", "urn:person:101"},
		// {"urn:movie:200", "cast", "urn:person:102"},
		// {"urn:movie:200", "cast", "urn:person:103"},
		// {"urn:movie:200", "sequel", "urn:movie:207"},

		// {"urn:movie:201", "title", "First Blood"},
		// {"urn:movie:201", "year", 1982},
		// {"urn:movie:201", "director", "urn:person:104"},
		// {"urn:movie:201", "cast", "urn:person:105"},
		// {"urn:movie:201", "cast", "urn:person:106"},
		// {"urn:movie:201", "cast", "urn:person:107"},
		// {"urn:movie:201", "sequel", "urn:movie:209"},

		// {"urn:movie:202", "title", "Predator"},
		// {"urn:movie:202", "year", 1987},
		// {"urn:movie:202", "director", "urn:person:108"},
		// {"urn:movie:202", "cast", "urn:person:101"},
		// {"urn:movie:202", "cast", "urn:person:109"},
		// {"urn:movie:202", "cast", "urn:person:110"},
		// {"urn:movie:202", "sequel", "urn:movie:211"},

		// {"urn:movie:203", "title", "Lethal Weapon"},
		// {"urn:movie:203", "year", 1987},
		// {"urn:movie:203", "director", "urn:person:111"},
		// {"urn:movie:203", "cast", "urn:person:112"},
		// {"urn:movie:203", "cast", "urn:person:113"},
		// {"urn:movie:203", "cast", "urn:person:114"},
		// {"urn:movie:203", "sequel", "urn:movie:212"},

		// {"urn:movie:204", "title", "RoboCop"},
		// {"urn:movie:204", "year", 1987},
		// {"urn:movie:204", "director", "urn:person:115"},
		// {"urn:movie:204", "cast", "urn:person:116"},
		// {"urn:movie:204", "cast", "urn:person:117"},
		// {"urn:movie:204", "cast", "urn:person:118"},

		// {"urn:movie:205", "title", "Commando"},
		// {"urn:movie:205", "year", 1985},
		// {"urn:movie:205", "director", "urn:person:119"},
		// {"urn:movie:205", "cast", "urn:person:101"},
		// {"urn:movie:205", "cast", "urn:person:120"},
		// {"urn:movie:205", "cast", "urn:person:121"},

		// {"urn:movie:206", "title", "Die Hard"},
		// {"urn:movie:206", "year", 1988},
		// {"urn:movie:206", "director", "urn:person:108"},
		// {"urn:movie:206", "cast", "urn:person:122"},
		// {"urn:movie:206", "cast", "urn:person:123"},
		// {"urn:movie:206", "cast", "urn:person:124"},

		// {"urn:movie:207", "title", "Terminator 2: Judgment Day"},
		// {"urn:movie:207", "year", 1991},
		// {"urn:movie:207", "director", "urn:person:100"},
		// {"urn:movie:207", "cast", "urn:person:101"},
		// {"urn:movie:207", "cast", "urn:person:102"},
		// {"urn:movie:207", "cast", "urn:person:125"},
		// {"urn:movie:207", "cast", "urn:person:126"},
		// {"urn:movie:207", "sequel", "urn:movie:208"},

		// {"urn:movie:208", "title", "Terminator 3: Rise of the Machines"},
		// {"urn:movie:208", "year", 2003},
		// {"urn:movie:208", "director", "urn:person:127"},
		// {"urn:movie:208", "cast", "urn:person:101"},
		// {"urn:movie:208", "cast", "urn:person:128"},
		// {"urn:movie:208", "cast", "urn:person:129"},

		// {"urn:movie:209", "title", "Rambo: First Blood Part II"},
		// {"urn:movie:209", "year", 1985},
		// {"urn:movie:209", "director", "urn:person:130"},
		// {"urn:movie:209", "cast", "urn:person:105"},
		// {"urn:movie:209", "cast", "urn:person:106"},
		// {"urn:movie:209", "cast", "urn:person:131"},
		// {"urn:movie:209", "sequel", "urn:movie:210"},

		// {"urn:movie:210", "title", "Rambo III"},
		// {"urn:movie:210", "year", 1988},
		// {"urn:movie:210", "director", "urn:person:132"},
		// {"urn:movie:210", "cast", "urn:person:105"},
		// {"urn:movie:210", "cast", "urn:person:106"},
		// {"urn:movie:210", "cast", "urn:person:133"},

		// {"urn:movie:211", "title", "Predator 2"},
		// {"urn:movie:211", "year", 1990},
		// {"urn:movie:211", "director", "urn:person:134"},
		// {"urn:movie:211", "cast", "urn:person:113"},
		// {"urn:movie:211", "cast", "urn:person:114"},
		// {"urn:movie:211", "cast", "urn:person:135"},

		// {"urn:movie:212", "title", "Lethal Weapon 2"},
		// {"urn:movie:212", "year", 1989},
		// {"urn:movie:212", "director", "urn:person:111"},
		// {"urn:movie:212", "cast", "urn:person:112"},
		// {"urn:movie:212", "cast", "urn:person:113"},
		// {"urn:movie:212", "cast", "urn:person:136"},
		// {"urn:movie:212", "sequel", "urn:movie:213"},

		// {"urn:movie:213", "title", "Lethal Weapon 3"},
		// {"urn:movie:213", "year", 1992},
		// {"urn:movie:213", "director", "urn:person:111"},
		// {"urn:movie:213", "cast", "urn:person:112"},
		// {"urn:movie:213", "cast", "urn:person:113"},
		// {"urn:movie:213", "cast", "urn:person:136"},

		// {"urn:movie:214", "title", "Alien"},
		// {"urn:movie:214", "year", 1979},
		// {"urn:movie:214", "director", "urn:person:137"},
		// {"urn:movie:214", "cast", "urn:person:138"},
		// {"urn:movie:214", "cast", "urn:person:139"},
		// {"urn:movie:214", "cast", "urn:person:140"},
		// {"urn:movie:214", "sequel", "urn:movie:215"},

		// {"urn:movie:215", "title", "Aliens"},
		// {"urn:movie:215", "year", 1986},
		// {"urn:movie:215", "director", "urn:person:100"},
		// {"urn:movie:215", "cast", "urn:person:139"},
		// {"urn:movie:215", "cast", "urn:person:141"},
		// {"urn:movie:215", "cast", "urn:person:103"},

		// {"urn:movie:216", "title", "Mad Max"},
		// {"urn:movie:216", "year", 1979},
		// {"urn:movie:216", "director", "urn:person:142"},
		// {"urn:movie:216", "cast", "urn:person:112"},
		// {"urn:movie:216", "cast", "urn:person:143"},
		// {"urn:movie:216", "cast", "urn:person:144"},
		// {"urn:movie:216", "sequel", "urn:movie:217"},

		// {"urn:movie:217", "title", "Mad Max 2"},
		// {"urn:movie:217", "year", 1981},
		// {"urn:movie:217", "director", "urn:person:142"},
		// {"urn:movie:217", "cast", "urn:person:112"},
		// {"urn:movie:217", "cast", "urn:person:145"},
		// {"urn:movie:217", "cast", "urn:person:146"},
		// {"urn:movie:217", "sequel", "urn:movie:218"},

		// {"urn:movie:218", "title", "Mad Max Beyond Thunderdome"},
		// {"urn:movie:218", "year", 1985},
		// {"urn:movie:218", "director", "urn:person:142"},
		// {"urn:movie:218", "director", "urn:person:147"},
		// {"urn:movie:218", "cast", "urn:person:112"},
		// {"urn:movie:218", "cast", "urn:person:148"},

		// {"urn:movie:219", "title", "Braveheart"},
		// {"urn:movie:219", "year", 1995},
		// {"urn:movie:219", "director", "urn:person:112"},
		// {"urn:movie:219", "cast", "urn:person:112"},
		// {"urn:movie:219", "cast", "urn:person:149"},
	}
}
