# Checkliste zu Goroutinen

* Wie lange soll die Goroutine laufen?
* Wie wird sie beendet?
* Welche Elemente blockieren den Programmfluss?
* Kann es zu einer Leaking Goroutine kommen?
* Was passiert auf der anderen Seite der Channels?

# Prüfung von Channels
## Channel Besonderheiten

* Deadlock
    * Daten an einen nil-Channel zu senden
    * Daten aus einen nil-Channel zu lesen
* Panik
    * Einen nil-Channel schließen
    * Einen bereits geschlossenen Channel zu schließen
* Nullwert
    * Daten aus einem geschlossenen Channel lesen
