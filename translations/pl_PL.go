package translations

func initEnglishTranslation() Translation {
	translation := createTranslation()

	translation.put("requires-js", "Ta strona wymaga Javascript aby działała poprawnie.")

	translation.put("start-the-game", "Rozpocznij Grę")
	translation.put("start", "Start")
	translation.put("game-not-started-title", "Gra się nie zaczęła")
	translation.put("waiting-for-host-to-start", "Poczekaj na hosta gry.")

	translation.put("round", "Runda")
	translation.put("toggle-soundeffects", "Efekty Dzwiękowe")
	translation.put("change-your-name", "Zmień swoją nazwę")
	translation.put("randomize", "Losuj")
	translation.put("apply", "Zaaplkikuj")
	translation.put("save", "Zapisz")
  translation.put("toggle-fullscreen", "Wielki Ekran (fullscreen)")
	translation.put("show-help", "Pokaż Pomoc")
	translation.put("votekick-a-player", "Zagłosuj aby wyrzucić gracza.")
	translation.put("time-left", "Czas")

	translation.put("last-turn", "(Ostatnia Runda: %s)")

	translation.put("drawer-kicked", "Związku z tym że gracz został wyrzucony nikt nie dostanie punktów!")
	translation.put("self-kicked", "Zostałeś wyrzucony")
	translation.put("kick-vote", "(%s/%s) zagłosował na wyrzucenie %s.")
	translation.put("player-kicked", "Gracz został wyrzucony.")
	translation.put("owner-change", "%s jest teraz właścicielem lobby.")

	translation.put("change-lobby-settings", "Zmień ustawienia lobby")
	translation.put("lobby-settings-changed", "Ustawienia Lobby zmienione")
	translation.put("advanced-settings", "Zaawansowane ustawienia")
	translation.put("word-language", "Słowo-Język")
	translation.put("drawing-time-setting", "Czas Rysowania")
	translation.put("rounds-setting", "Rundy")
	translation.put("max-players-setting", "Maksimum Graczy")
	translation.put("public-lobby-setting", "Publiczne Lobby")
	translation.put("custom-words", "Własne Wyrazy")
	translation.put("custom-words-info", "Wpisz tutaj swoje własne wyrazy, przedziel je przecinkami")
	translation.put("custom-words-chance-setting", "Szansa na Własny Wyraz")
	translation.put("players-per-ip-limit-setting", "Limit graczy na jedno IP")
	translation.put("enable-votekick-setting", "Pozwól na VoteKick")
	translation.put("save-settings", "Zapisz ustawienia")
	translation.put("input-contains-invalid-data", "Twój wpis miał niepoprawne dane:")
	translation.put("please-fix-invalid-input", "Popraw to i napisz ponownie.")
	translation.put("create-lobby", "Stwórz Lobby")

	translation.put("players", "Gracze")
	translation.put("refresh", "Odśwież")
	translation.put("join-lobby", "Dołącz do Lobby")

	translation.put("message-input-placeholder", "Wpisuj tutaj swoje rozmyślenia i wiadomośći")

	translation.put("choose-a-word", "Wybierz słowo")
	translation.put("waiting-for-word-selection", "Oczekiwanie na wybranie słowa")
	//This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "wybiera słowo.")

	translation.put("close-guess", "'%s' jest bardzo blisko.")
	translation.put("correct-guess", "Zgadłeś Słowo.")
	translation.put("correct-guess-other-player", "%s dobrze zgadł/a hasło.")
	translation.put("round-over", "Zamiania, żadne hasło nie zostało wybrane.")
	translation.put("round-over-no-word", "No szkoda, hasło to '%s'.")
	translation.put("game-over-win", "Brawo Wygrałeś!")
	translation.put("game-over", "Zająłeś %s. z %s punktami")

	translation.put("change-active-color", "Zmień swój kolor")
	translation.put("use-pencil", "Użyj ołówka")
	translation.put("use-eraser", "Użyj gumki")
	translation.put("use-fill-bucket", "Użyj wiadra")
	translation.put("change-pencil-size-to", "Zmień rozmiar ołowka/gumki na %s")
	translation.put("clear-canvas", "Wyczyść obraz")

	translation.put("connection-lost", "Zerwano połączenie!")
	translation.put("connection-lost-text", "Próba połączenia"+
		" ...\n\nsprawdź czy ci internet działa bambiku.\nJeśli nadal masz problem to skontaktuj pana Informatyka")
	translation.put("error-connecting", "Error connecting to server")
	translation.put("error-connecting-text",
		"Scribble.rs couldn't establish a socket connection.\n\nWhile your internet "+
			"connection seems to be working, either the\nserver or your firewall hasn't "+
			"been configured correctly.\n\nTo retry, reload the page.")

	translation.put("message-too-long", "Wiadomość za długa.")

	//Help dialog
	translation.put("controls", "Sterowanie")
	translation.put("pencil", "Ołówek")
	translation.put("eraser", "Gumka")
	translation.put("fill-bucket", "Wiadro")
	translation.put("switch-tools-intro", "Możesz zmieniać narzędzia skrótami")
	translation.put("switch-pencil-sizes", "Możesz zmieniać rozmiar ołowka %s na %s.")

	//Generic words
	//"close" as in "closing the window"
	translation.put("close", "Zamknij")
	translation.put("no", "Nie")
	translation.put("yes", "Tak")
	translation.put("system", "System")

	translation.put("source-code", "Kod Źródłowy")
	translation.put("help", "Pomoc")
	translation.put("contact", "Kontakt")
	translation.put("submit-feedback", "Recenzja")
	translation.put("stats", "Status")

	RegisterTranslation("en", translation)
	RegisterTranslation("en-pl", translation)

	return translation
}
