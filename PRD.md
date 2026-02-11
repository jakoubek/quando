# Product Requirements Document: quando

## 1. Projektübersicht

### Ziel und Vision
`quando` ist eine eigenständige Go-Bibliothek für intuitive und idiomatische Datumsberechnungen. Sie bietet eine Fluent API für komplexe Datums-Operationen, die mit der Go-Standardbibliothek umständlich oder unmöglich sind. Die Library ist der Kern der geplanten DatesAPI v2 und soll als wiederverwendbare Komponente in beliebigen Go-Projekten einsetzbar sein.

**Vision:** Die bevorzugte Go-Library für Datumsberechnungen – so natürlich und intuitiv wie Moment.js oder Carbon, aber Go-idiomatisch und ohne externe Dependencies.

### Zielgruppe
- **Primär:** Go-Entwickler, die komplexe Datumsberechnungen durchführen müssen
- **Sekundär:** Entwickler der DatesAPI v2 (interner First User)
- **Technisches Level:** Erfahrene Go-Entwickler, die Go-Idiome und `time.Time` kennen

### Erfolgs-Metriken
- **Adoption:** Verwendung in der DatesAPI v2 als Proof-of-Concept
- **Code-Reduktion:** Typische Datums-Operationen in 1 Zeile statt 5-10 Zeilen
- **Test-Coverage:** Minimum 95% für alle Kalkulationsfunktionen
- **Performance:** Alle Operationen unter 1µs (außer Parsing)
- **Zero Dependencies:** Ausschließlich Go stdlib (außer optionale i18n-Erweiterungen)

### Projektscope

**Phase 1 – In Scope:**
- Datums-Arithmetik (Add/Sub mit allen Zeiteinheiten)
- Snap-to/Ankerpunkte (StartOf, EndOf, Next, Prev)
- Differenz-Berechnung (mit Human-Format)
- Datums-Inspektion (WeekNumber, Quarter, DayOfYear, etc.)
- Formatierung (Presets + Custom Layouts + i18n)
- Zeitzone-Support (Konvertierung + DST-Handling)
- Unix-Timestamp-Konvertierung
- Parsing (automatisch + explizit + relativ)

**Out of Scope (spätere Phasen):**
- HTTP/API-Schicht (separater Webserver)
- Feiertage & Arbeitstage (Phase 3)
- Datums-Serien/Ranges (Phase 2)
- Batch-Operationen (Phase 2)

---

## 2. Funktionale Anforderungen

### Kern-Features

#### 2.1 Datums-Arithmetik (Must-have)

Verkettbare Add/Sub-Operationen für alle Zeiteinheiten.

**API:**
```go
quando.From(time.Now()).Add(2, quando.Days)
quando.From(time.Now()).Add(2, quando.Months).Sub(3, quando.Days)
quando.Now().Add(1, quando.Years)
```

**Unterstützte Einheiten:** `Seconds`, `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Quarters`, `Years`

**Edge Cases:**
- **Monatsende-Overflow:** Bei Addition von Monaten wird auf das Monatsende gekürzt, falls das Zieldatum nicht existiert
  - `2026-01-31` + 1 Monat = `2026-02-28` (Monatsende Februar)
  - `2026-01-24` + 1 Monat = `2026-02-24` (regulär)
  - `2026-05-31` + 1 Monat = `2026-06-30` (Juni hat nur 30 Tage)

#### 2.2 Snap-to / Ankerpunkte (Must-have)

Sprung zum Anfang/Ende einer Zeiteinheit oder zum nächsten/vorherigen Wochentag.

**API:**
```go
quando.From(date).StartOf(quando.Month)   // 1. des Monats, 00:00:00
quando.From(date).EndOf(quando.Quarter)   // Letzter Tag des Quartals, 23:59:59
quando.From(date).Next(time.Monday)       // Nächster Montag (nie heute)
quando.From(date).Prev(time.Friday)       // Vorheriger Freitag (nie heute)
```

**Snap-Einheiten:** `Week`, `Month`, `Quarter`, `Year`

**Verhalten:**
- **StartOf(Week):** Montag 00:00:00 (ISO 8601 Default, konfigurierbar)
- **EndOf(Week):** Sonntag 23:59:59 bei Montag-Start, Samstag 23:59:59 bei Sonntag-Start
- **Next(Weekday):** Immer der NÄCHSTE Wochentag (nie heute, auch wenn heute der gleiche Wochentag ist)
- **Prev(Weekday):** Immer der VORHERIGE Wochentag (nie heute)

**Quartals-Definition:**
- Q1 = Januar–März
- Q2 = April–Juni
- Q3 = Juli–September
- Q4 = Oktober–Dezember

Keine Konfiguration von Geschäftsjahren in Phase 1.

#### 2.3 Differenz-Berechnung (Must-have)

Berechnung der Differenz zwischen zwei Daten in verschiedenen Einheiten.

**API:**
```go
d := quando.Diff(date1, date2)
d.Days()         // int: 319
d.Weeks()        // int: 45
d.Months()       // int: 10 (abgerundet)
d.MonthsFloat()  // float64: 10.516... (präzise)
d.Years()        // int: 0
d.YearsFloat()   // float64: 0.876...
d.Human()        // "10 months, 16 days" (adaptive Granularität)
d.Human(quando.LangDE) // "10 Monate, 16 Tage"
```

**Präzision:**
- Integer-Varianten (`Months()`, `Years()`) geben abgerundete Werte zurück
- Float-Varianten (`MonthsFloat()`, `YearsFloat()`) für präzise Berechnungen

**Human-Format:**
Adaptive Granularität – immer die zwei größten relevanten Einheiten:

| Differenz             | Ausgabe (EN)          | Ausgabe (DE)           |
|-----------------------|-----------------------|------------------------|
| 10 Monate, 16 Tage    | "10 months, 16 days"  | "10 Monate, 16 Tage"   |
| 2 Tage, 5 Stunden     | "2 days, 5 hours"     | "2 Tage, 5 Stunden"    |
| 3 Stunden, 20 Minuten | "3 hours, 20 minutes" | "3 Stunden, 20 Minuten"|
| 45 Sekunden           | "45 seconds"          | "45 Sekunden"          |
| 0                     | "0 seconds"           | "0 Sekunden"           |

#### 2.4 Datums-Inspektion (Must-have)

Abfrage von Meta-Informationen zu einem Datum.

**API (aggregiert):**
```go
info := quando.From(date).Info()
info.WeekNumber   // int: ISO 8601 Week Number
info.Quarter      // int: 1–4
info.DayOfYear    // int: 1–366
info.IsWeekend    // bool: Samstag oder Sonntag
info.IsLeapYear   // bool: Schaltjahr
info.Unix         // int64: Unix Timestamp
```

**API (einzeln):**
```go
quando.From(date).WeekNumber()  // 7
quando.From(date).Quarter()     // 1
quando.From(date).DayOfYear()   // 40
quando.From(date).IsWeekend()   // true
quando.From(date).IsLeapYear()  // false
quando.From(date).Unix()        // 1770595200
```

**Konventionen:**
- **WeekNumber:** ISO 8601 (Montag = erster Tag, Woche 1 = erste Woche mit Donnerstag)
- **IsWeekend:** Samstag + Sonntag (nicht konfigurierbar in Phase 1)
- **IsLeapYear:** Standard-Regel (durch 4 teilbar, außer Jahrhundert, außer durch 400 teilbar)

#### 2.5 Formatierung (Must-have)

Erweiterte Formatierung mit Presets, Custom Layouts und Mehrsprachigkeit.

**Preset-Formate:**
```go
quando.From(date).Format(quando.ISO)      // "2026-02-09"
quando.From(date).Format(quando.EU)       // "09.02.2026"
quando.From(date).Format(quando.US)       // "02/09/2026"
quando.From(date).Format(quando.Long)     // "February 9, 2026"
quando.From(date).Format(quando.RFC2822)  // "Mon, 09 Feb 2026 00:00:00 +0000"
```

**Custom Layouts (Go-Standard):**
```go
quando.From(date).Format("Monday, 2. January 2006")  // "Monday, 9. February 2026"
```

**Mehrsprachigkeit:**
```go
quando.From(date).Lang(quando.DE).Format(quando.Long)
// "9. Februar 2026"

quando.From(date).Lang(quando.DE).Format("Monday, 2. January 2006")
// "Montag, 9. Februar 2026"
```

**Sprachregeln:**
- Nur `Long` und Custom Layouts sind sprachabhängig
- ISO, EU, US, RFC2822 sind immer sprachunabhängig
- **Phase 1 Sprachen:** EN (Default), DE (Must-have)
- **Spätere Phasen:** Die 21 weiteren Sprachen aus v1 (ES, FA, FR, HI, ID, IT, JP, KR, MS_MY, NL, PL, PT, RO, RU, SE, TH, TR, UK, VI, ZH_CN, ZH_TW)

#### 2.6 Zeitzone-Support (Must-have)

Konvertierung zwischen Zeitzonen mit korrektem DST-Handling.

**API:**
```go
// Datum in Zeitzone
quando.From(time.Now()).In("Europe/Berlin")

// Konvertierung
quando.From(date).In("America/New_York")
```

**Default-Verhalten:**
- **Default-Zeitzone:** UTC (wenn nicht explizit gesetzt)
- **DST-Handling:** `Add(1, Days)` bedeutet "selbe Uhrzeit am nächsten Kalendertag", NICHT 24 Stunden
  - Beispiel: `2026-03-31 02:00` (CET) + 1 Day = `2026-04-01 02:00` (CEST), auch wenn dies nur 23 Stunden sind
  - Rationale: Menschen denken in Kalendertagen, nicht in Stunden-Deltas

#### 2.7 Unix-Timestamp-Konvertierung (Must-have)

Bidirektionale Konvertierung zwischen `time.Time` und Unix-Timestamps.

**API:**
```go
// time.Time → Unix
quando.From(date).Unix()  // 1770595200

// Unix → time.Time
quando.FromUnix(1770595200)  // Date
```

**Unterstützung:**
- Positive und negative Timestamps (vor 1970) werden unterstützt
- Keine künstlichen Datumsbeschränkungen – Go's `time.Time` Range (Jahr 0–9999+)

#### 2.8 Parsing (Must-have)

Automatisches und explizites Parsing von Datums-Strings.

**Automatisches Parsing:**
```go
quando.Parse("2026-02-09")       // ISO
quando.Parse("09.02.2026")       // EU (Punkt-Trennzeichen)
quando.Parse("2026/02/09")       // ISO mit Slash
quando.Parse("Mon, 09 Feb 2026") // RFC2822
```

**Mehrdeutigkeits-Regel:**
Slash-Formate ohne Jahr-Prefix sind mehrdeutig und führen zu einem Error:

| Eingabe      | Erkennung         | Begründung                         |
|--------------|-------------------|------------------------------------|
| `2026-02-01` | ✅ ISO, eindeutig  | Standard-Format                    |
| `01.02.2026` | ✅ EU, eindeutig   | Punkt = EU-Konvention              |
| `2026/02/09` | ✅ ISO, eindeutig  | Jahr-Prefix ist eindeutig          |
| `01/02/2026` | ❌ Error           | Mehrdeutig (US vs. EU)             |

**Explizites Parsing:**
```go
quando.ParseWithLayout("01/02/2026", "02/01/2006")  // EU-Format
quando.ParseWithLayout("01/02/2026", "01/02/2006")  // US-Format
```

**Relative Ausdrücke (Must-have in Phase 1):**
```go
quando.ParseRelative("today")       // Heute 00:00:00
quando.ParseRelative("tomorrow")    // Morgen 00:00:00
quando.ParseRelative("yesterday")   // Gestern 00:00:00
quando.ParseRelative("+2 days")     // Heute + 2 Tage
quando.ParseRelative("-1 week")     // Heute - 1 Woche
quando.ParseRelative("+3 months")   // Heute + 3 Monate
```

Komplexere Ausdrücke (`"next monday"`, `"start of month"`) sind Nice-to-have für spätere Versionen.

---

### User Stories

**Als Go-Entwickler möchte ich...**

1. **Komplexe Arithmetik:** ...verkettete Datums-Operationen in einer Zeile schreiben können, damit mein Code lesbar bleibt
   - Akzeptanzkriterium: `quando.Now().Add(2, Months).Sub(3, Days)` funktioniert

2. **Monatsende-Arithmetik:** ...Monate addieren ohne manuell Overflow-Fälle zu behandeln
   - Akzeptanzkriterium: `31. Jan + 1 Monat = 28. Feb` (automatisch)

3. **Quartalsberechnungen:** ...zum Quartalsanfang/-ende springen können
   - Akzeptanzkriterium: `quando.Now().StartOf(Quarter)` gibt 1. Jan/Apr/Jul/Okt zurück

4. **Differenz-Formatierung:** ...Datumsdifferenzen menschenlesbar ausgeben
   - Akzeptanzkriterium: `Diff(a, b).Human()` gibt "10 months, 16 days" zurück

5. **Mehrsprachigkeit:** ...Datums-Strings in verschiedenen Sprachen formatieren
   - Akzeptanzkriterium: `.Lang(DE).Format(Long)` gibt "9. Februar 2026" zurück

6. **Zeitzone-Transparenz:** ...Datums-Arithmetik über DST-Umstellungen hinweg korrekt durchführen
   - Akzeptanzkriterium: `Add(1, Days)` bedeutet "nächster Kalendertag", nicht "24 Stunden"

7. **Testbarkeit:** ...deterministische Tests schreiben können
   - Akzeptanzkriterium: `quando.NewClock(fixedTime)` für Test-Fixtures

---

### Detaillierte Workflows

#### Workflow 1: Datums-Arithmetik mit Verkettung

```go
// Szenario: Berechne "Letzter Tag des übernächsten Quartals"
date := quando.Now().
    Add(2, quando.Quarters).
    EndOf(quando.Quarter)

// Szenario: "Erster Montag nach Monatsende"
date := quando.Now().
    EndOf(quando.Month).
    Next(time.Monday)
```

#### Workflow 2: Differenz-Berechnung mit Formatierung

```go
start := quando.MustParse("2025-01-15")
end := quando.Now()

diff := quando.Diff(start.Time(), end.Time())

fmt.Printf("Tage: %d\n", diff.Days())
fmt.Printf("Monate: %d\n", diff.Months())
fmt.Printf("Lesbar: %s\n", diff.Human(quando.LangDE))
```

#### Workflow 3: Parsing unbekannter Formate

```go
// Automatisches Parsing
date, err := quando.Parse(userInput)
if err != nil {
    // Fallback: Explizites Format
    date, err = quando.ParseWithLayout(userInput, "02/01/2006")
}
```

#### Workflow 4: Zeitzone-Konvertierung

```go
// UTC → Berlin
utcDate := quando.FromUnix(1770595200)
berlinDate := utcDate.In("Europe/Berlin")

// Arithmetik in spezifischer Zeitzone
date := quando.Now().
    In("America/New_York").
    Add(1, quando.Days)
```

---

### Feature-Prioritäten

| Feature                     | Priorität  | Rationale                                  |
|-----------------------------|------------|--------------------------------------------|
| Datums-Arithmetik           | Must-have  | Kern-Feature, primärer Use Case            |
| Snap-to/Ankerpunkte         | Must-have  | Häufiger Use Case, schwer mit stdlib      |
| Differenz-Berechnung (int)  | Must-have  | Häufiger Use Case                          |
| Differenz-Berechnung (float)| Must-have  | Für präzise Berechnungen notwendig         |
| Human-Format (EN, DE)       | Must-have  | Differenzierung zu anderen Libraries       |
| Parsing (automatisch)       | Must-have  | Eingangs-Punkt für User Input              |
| Zeitzone-Support            | Must-have  | Essentiell für korrekte Berechnungen       |
| Datums-Inspektion           | Must-have  | Convenience, wenig Aufwand                 |
| Unix-Timestamp              | Must-have  | Standard-Interop mit APIs                  |
| Formatierung (Presets)      | Must-have  | Häufiger Use Case                          |
| ParseRelative (basic)       | Must-have  | User-freundlich, von v1 gewohnt            |
| Formatierung (Custom)       | Must-have  | Flexibilität                               |
| ParseRelative (advanced)    | Nice-to-have| Komplexität vs. Nutzen                   |
| Weitere Sprachen (21)       | Nice-to-have| Internationale Nutzung, aber aufwändig   |

---

## 3. Technische Anforderungen

### Performance-Ziele

- **Arithmetik-Operationen:** < 1 µs pro Operation (Add, Sub, StartOf, EndOf)
- **Differenz-Berechnung:** < 1 µs für Integer-Varianten, < 2 µs für Float-Varianten
- **Formatierung:** < 5 µs ohne i18n, < 10 µs mit i18n
- **Parsing (automatisch):** < 10 µs für eindeutige Formate
- **Parsing (relativ):** < 20 µs
- **Memory Allocations:** Keine Allocations bei verketteten Operationen (außer final result)

### Concurrent User-Kapazität

Nicht anwendbar – reine Library ohne Server-Komponente.

### Real-time Features

Nicht anwendbar – keine Echtzeit-Features, keine WebSockets/SSE.

### Sicherheitsstandards

- **Input Validation:** Alle Parse-Funktionen müssen ungültige Eingaben sicher mit Error zurückweisen (NIEMALS panic)
- **Overflow-Schutz:** Datums-Arithmetik muss Overflow-Szenarien handhaben (z.B. Jahr > 9999)
- **Timezone-Sicherheit:** IANA-Timezone-Namen validieren, bei ungültigen Namen Error zurückgeben

### Compliance-Vorgaben

- **ISO 8601:** Compliance für Datumsformate, Wochennummern, Zeitzonen
- **IANA Timezone Database:** Verwendung der Standard-Zeitzonen-Datenbank

### Plattform-Support

- **Go-Version:** Minimum Go 1.22+ (für aktuelle stdlib-Features)
- **Betriebssysteme:** Alle von Go unterstützten Plattformen (Linux, macOS, Windows, BSD)
- **Architekturen:** Alle von Go unterstützten Architekturen (amd64, arm64, 386, arm)

---

## 4. Datenarchitektur

Nicht anwendbar – keine Datenbank, keine Persistierung. Alle Daten sind transient (in-memory).

---

## 5. API & Interface-Spezifikation

### Go-API-Design

#### Haupt-Typ: `quando.Date`

```go
// Date ist der zentrale Typ der Library
// Er wrapped time.Time und bietet eine Fluent API
type Date struct {
    t    time.Time
    lang Lang  // optional, für Formatierung
}

// Konvertierung time.Time → quando.Date
func From(t time.Time) Date

// Konvertierung quando.Date → time.Time
func (d Date) Time() time.Time
```

**Design-Rationale:**
- Eigener Wrapper-Typ notwendig für Fluent API (Verkettung)
- Kein Reimplementieren von `time.Time` – Delegation an stdlib
- Einfache bidirektionale Konvertierung

#### Einheiten-Konstanten

```go
type Unit int

const (
    Seconds Unit = iota
    Minutes
    Hours
    Days
    Weeks
    Months
    Quarters
    Years
)
```

**Design-Rationale:**
- Typsichere Konstanten statt Strings (Compile-Time Safety)
- `iota` für klare Ordnung
- Interne `ParseUnit(string) Unit` für externe Eingaben (API-Layer, ParseRelative)

#### Sprach-Konstanten

```go
type Lang string

const (
    LangEN Lang = "en"  // English (Default)
    LangDE Lang = "de"  // Deutsch
    // Weitere 21 Sprachen in späteren Versionen
)
```

#### Fehlerbehandlung

Alle Funktionen, die fehlschlagen können, geben `(Result, error)` zurück:

```go
func Parse(s string) (Date, error)
func ParseWithLayout(s, layout string) (Date, error)
func ParseRelative(s string) (Date, error)
func FromUnix(sec int64) (Date, error)  // kann bei Overflow fehlschlagen
```

**Convenience-Varianten (Panic bei Fehler):**
```go
func MustParse(s string) Date  // nur für Tests/Init
```

**Design-Regel:** Library darf NIE panicken (außer Must-Varianten).

#### Clock-Abstraktion (für Testbarkeit)

```go
// Clock ermöglicht Time-Injection für Tests
type Clock interface {
    Now() Date
    From(t time.Time) Date
}

// DefaultClock verwendet time.Now()
func NewClock() Clock

// FixedClock für deterministische Tests
func NewFixedClock(t time.Time) Clock
```

**Verwendung:**
```go
// Produktion
date := quando.Now()  // verwendet DefaultClock

// Tests
clock := quando.NewFixedClock(time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC))
date := clock.Now()
date2 := clock.From(otherTime).Add(2, quando.Days)
```

### Öffentliche API (Package-Level Funktionen)

```go
// Konstruktoren
func Now() Date
func From(t time.Time) Date
func FromUnix(sec int64) (Date, error)

// Parsing
func Parse(s string) (Date, error)
func ParseWithLayout(s, layout string) (Date, error)
func ParseRelative(s string) (Date, error)
func MustParse(s string) Date

// Differenz
func Diff(a, b time.Time) Duration

// Clock-Factory (für Tests)
func NewClock() Clock
func NewFixedClock(t time.Time) Clock
```

### Methoden auf `quando.Date`

```go
// Arithmetik
func (d Date) Add(value int, unit Unit) Date
func (d Date) Sub(value int, unit Unit) Date

// Snap-to
func (d Date) StartOf(unit Unit) Date
func (d Date) EndOf(unit Unit) Date
func (d Date) Next(weekday time.Weekday) Date
func (d Date) Prev(weekday time.Weekday) Date

// Inspektion
func (d Date) Info() DateInfo
func (d Date) WeekNumber() int
func (d Date) Quarter() int
func (d Date) DayOfYear() int
func (d Date) IsWeekend() bool
func (d Date) IsLeapYear() bool
func (d Date) Unix() int64

// Formatierung
func (d Date) Format(format Format) string
func (d Date) FormatLayout(layout string) string
func (d Date) Lang(lang Lang) Date  // Fluent API

// Zeitzone
func (d Date) In(location string) (Date, error)

// Konvertierung
func (d Date) Time() time.Time
```

### Methoden auf `quando.Duration` (Differenz)

```go
type Duration struct {
    // private fields
}

func (dur Duration) Seconds() int64
func (dur Duration) Minutes() int64
func (dur Duration) Hours() int64
func (dur Duration) Days() int
func (dur Duration) Weeks() int
func (dur Duration) Months() int
func (dur Duration) MonthsFloat() float64
func (dur Duration) Years() int
func (dur Duration) YearsFloat() float64
func (dur Duration) Human() string
func (dur Duration) Human(lang Lang) string
```

### Format-Konstanten

```go
type Format int

const (
    ISO Format = iota     // "2026-02-09"
    EU                    // "09.02.2026"
    US                    // "02/09/2026"
    Long                  // "February 9, 2026" (sprachabhängig)
    RFC2822               // "Mon, 09 Feb 2026 00:00:00 +0000"
)
```

---

## 6. Benutzeroberfläche

Nicht anwendbar – reine Go-Library ohne UI-Komponente.

---

## 7. Nicht-funktionale Anforderungen

### Verfügbarkeit und Uptime

Nicht anwendbar – keine Server-Komponente.

### Graceful Shutdown und Signal-Handling

Nicht anwendbar – keine Server-Komponente.

### Backup- und Recovery

Nicht anwendbar – keine Datenpersistierung.

### Monitoring und Observability

Nicht anwendbar – Library-Nutzer sind selbst verantwortlich für Monitoring.

**Empfehlung für Library-Nutzer:**
- Performance-kritische Operationen mit Benchmarks messen
- Error-Rates von Parse-Funktionen loggen

### Logging-Strategie

Keine interne Logging-Komponente. Library gibt Errors über Return-Values zurück, niemals über Logging.

**Design-Rationale:** Libraries sollten keine Logs schreiben – das ist Aufgabe der Applikation.

### Deployment

#### Repository-Struktur

```
quando/
├── quando.go           # Haupt-API
├── date.go             # Date-Type und Core-Methoden
├── arithmetic.go       # Add, Sub
├── snap.go             # StartOf, EndOf, Next, Prev
├── diff.go             # Differenz-Berechnung
├── inspect.go          # WeekNumber, Quarter, etc.
├── format.go           # Formatierung
├── parse.go            # Parsing
├── clock.go            # Clock-Abstraktion
├── i18n.go             # Internationalisierung (EN, DE)
├── i18n_test.go
├── example_test.go     # Godoc-Examples
├── quando_test.go      # Unit-Tests
├── bench_test.go       # Benchmarks
├── go.mod
├── go.sum
├── README.md
├── LICENSE (MIT)
└── .github/
    └── workflows/
        └── ci.yml      # GitHub Actions
```

#### Lizenz

MIT License – Open Source geplant.

#### Versionierung

Semantic Versioning (semver):
- `v0.x.x` während Phase 1 (API nicht stabil)
- `v1.0.0` nach erfolgreichem Einsatz in DatesAPI v2
- Breaking Changes nur bei Major-Versions

### Skalierung und Load Balancing

Nicht anwendbar – keine Server-Komponente.

**Performance-Überlegungen:**
- Library ist thread-safe (alle Operationen auf unveränderlichen Daten)
- Keine Shared State – parallele Nutzung ohne Locks möglich
- Geeignet für hochparallelisierte Workloads (Goroutines)

---

## 8. Qualitätssicherung

### Definition of Done

Ein Feature ist "Done", wenn:

1. **Implementierung:**
   - Code folgt Go-Idiomen (go fmt, go vet, golangci-lint)
   - Alle exportierten Funktionen/Types dokumentiert (Godoc)
   - Fehlerbehandlung mit `error`-Return-Values (niemals panic)

2. **Tests:**
   - Unit-Tests für alle Funktionen (min. 95% Coverage)
   - Edge-Case-Tests (Schaltjahre, Monatsende, DST-Umstellungen)
   - Benchmarks für Performance-kritische Funktionen
   - Example-Tests für Godoc (in `example_test.go`)

3. **Dokumentierung:**
   - Godoc-Kommentare für alle Public APIs
   - README mit Code-Beispielen aktualisiert
   - Changelog-Eintrag

4. **Review:**
   - Code-Review abgeschlossen
   - CI/CD-Pipeline erfolgreich (Tests, Linting)

### Test-Anforderungen

#### Unit-Tests

**Mindest-Coverage:** 95% für alle Kalkulationsfunktionen

**Kritische Test-Szenarien:**

1. **Datums-Arithmetik:**
   - Monatsende-Overflow (31. Jan + 1 Monat = 28. Feb)
   - Schaltjahr-Handling (29. Feb in Schaltjahren)
   - Negative Arithmetik (Subtraktion über Jahresgrenzen)
   - Verkettung (mehrere Add/Sub-Operationen)

2. **Snap-to/Ankerpunkte:**
   - StartOf/EndOf für alle Einheiten (Week, Month, Quarter, Year)
   - Next/Prev bei gleichem Wochentag (muss überspringen)
   - EndOf(Week) mit verschiedenen Wochenbeginn-Einstellungen

3. **Differenz-Berechnung:**
   - Differenz über Jahresgrenzen
   - Differenz über Schaltjahre
   - Negative Differenzen (date1 < date2)
   - Human-Format mit verschiedenen Granularitäten

4. **Parsing:**
   - Alle unterstützten Formate (ISO, EU, RFC2822)
   - Ungültige Eingaben (Error-Handling)
   - Mehrdeutige Formate (Slash ohne Jahr-Prefix)
   - Relative Ausdrücke (today, +2 days, etc.)

5. **Zeitzone & DST:**
   - Konvertierung zwischen Zeitzonen
   - DST-Umstellung (Add(1, Days) über DST-Grenze)
   - Ungültige Timezone-Namen (Error-Handling)

6. **Datums-Inspektion:**
   - WeekNumber für ISO 8601 (Woche 1 = erste Woche mit Donnerstag)
   - Quarter-Berechnung für Grenzfälle (31. März, 30. Juni, etc.)
   - IsLeapYear für alle Regeln (durch 4, außer Jahrhundert, außer durch 400)

#### Benchmarks

Benchmarks für alle Performance-kritischen Funktionen:

```go
func BenchmarkAdd(b *testing.B)
func BenchmarkDiff(b *testing.B)
func BenchmarkParse(b *testing.B)
func BenchmarkFormat(b *testing.B)
```

**Performance-Ziele:**
- Add/Sub: < 1 µs
- Diff: < 2 µs
- Parse: < 10 µs
- Format: < 10 µs

#### Integration mit Testcontainers

Nicht anwendbar – keine Datenbank/externe Services.

### Launch-Kriterien

**Phase 1 Launch (v0.1.0):**
- Alle Must-have Features implementiert
- Test-Coverage > 95%
- Benchmarks erfüllen Performance-Ziele
- README mit Beispielen vollständig
- CI/CD-Pipeline funktioniert
- Code-Review abgeschlossen

**Production-Ready (v1.0.0):**
- Erfolgreich in DatesAPI v2 integriert
- Keine kritischen Bugs im Production-Einsatz (4+ Wochen)
- API-Stabilität erreicht (keine Breaking Changes mehr geplant)
- Umfangreiche Dokumentation (README, Godoc, Examples)

### Abnahme-Prozess

1. **Selbst-Review:** Entwickler prüft eigenen Code gegen DoD
2. **Code-Review:** Mindestens ein Review durch anderen Go-Entwickler
3. **CI/CD:** Alle automatisierten Tests und Lints bestanden
4. **Integration-Test:** Verwendung in DatesAPI v2 (Smoke-Test)
5. **Abnahme:** Product Owner prüft Feature gegen Acceptance Criteria

---

## 9. Technische Implementierungshinweise

### Go-Projektstruktur und Package-Layout

**Flat Package-Struktur:**

```
quando/
├── quando.go       # Package-Level Funktionen (Now, From, Parse, Diff)
├── date.go         # Date-Type, Core-Methoden, Konvertierung
├── arithmetic.go   # Add, Sub (Logik für Monatsende-Overflow)
├── snap.go         # StartOf, EndOf, Next, Prev
├── diff.go         # Duration-Type, Differenz-Berechnung
├── inspect.go      # WeekNumber, Quarter, DayOfYear, etc.
├── format.go       # Format, FormatLayout, Preset-Konstanten
├── parse.go        # Parse, ParseWithLayout, ParseRelative
├── clock.go        # Clock-Interface, DefaultClock, FixedClock
├── i18n.go         # Internationalisierung (EN, DE)
├── errors.go       # Custom Error-Types
└── internal/
    └── calc/       # Interne Hilfs-Funktionen (nicht exportiert)
```

**Design-Rationale:**
- **Flat Package:** Alle Funktionen unter `quando.*` verfügbar – einfache Benutzung
- **Keine Sub-Packages:** Vermeidet zyklische Dependencies und komplexe Imports
- **internal/:** Für nicht-exportierte Hilfs-Logik (z.B. Schaltjahr-Berechnung, Kalender-Arithmetik)
- **Ausnahme:** Wenn i18n-Daten sehr groß werden (22 Sprachen), optionales `quando/lang` Sub-Package für späteren Import

### Concurrency-Patterns

**Thread-Safety:**
- Alle `quando.Date`-Operationen sind immutable – jede Operation gibt ein neues `Date` zurück
- Keine Shared State, keine Mutexes notwendig
- Goroutine-safe ohne zusätzlichen Aufwand

**Design-Rationale:**
- Immutability vermeidet Data-Races und macht Library goroutine-safe
- Funktionaler Stil (Fluent API) begünstigt Immutability
- Performance-Overhead minimal (Stack-Allocation für kleine Structs)

**Nicht verwenden:**
- Keine Goroutines innerhalb der Library (Library-Code sollte synchron sein)
- Keine Channels, keine `errgroup` – das ist Aufgabe der Applikation

### Error-Handling-Strategie

**Prinzipien:**
1. **Niemals panic (außer Must-Varianten):** Alle Errors über Return-Values
2. **Sentinel Errors für bekannte Fehler:**
   ```go
   var (
       ErrInvalidFormat  = errors.New("invalid date format")
       ErrInvalidTimezone = errors.New("invalid timezone")
       ErrOverflow       = errors.New("date overflow")
   )
   ```
3. **Wrapped Errors für Kontext:**
   ```go
   return Date{}, fmt.Errorf("parsing date: %w", err)
   ```

**Error-Kategorien:**
- **Parse-Errors:** Ungültige Formate, mehrdeutige Eingaben
- **Timezone-Errors:** Unbekannte IANA-Namen
- **Overflow-Errors:** Datums-Arithmetik außerhalb des Go-Range

### Dependency Injection und Konfiguration

**Clock-Pattern (für Tests):**

```go
// Production-Code
date := quando.Now()

// Test-Code
clock := quando.NewFixedClock(time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC))
date := clock.Now()
```

**Konfiguration für Wochenbeginn (optional):**

```go
// Default: Montag (ISO 8601)
date.StartOf(quando.Week)

// Optional: Konfigurierbarer Wochenbeginn (Nice-to-have)
cfg := quando.Config{WeekStartDay: time.Sunday}
date := quando.WithConfig(cfg).From(time.Now()).StartOf(quando.Week)
```

**Design-Entscheidung:** Konfiguration ist Nice-to-have für Phase 1. Default-Verhalten (ISO 8601, UTC) sollte für 95% der Use Cases ausreichen.

### PostgreSQL Connection Pooling

Nicht anwendbar – keine Datenbank.

### Entwicklungs-Prioritäten

**Phase 1 Reihenfolge:**

1. **Woche 1-2: Core-Infrastruktur**
   - `Date`-Type, `From()`, `Time()`, `Unix()`, `FromUnix()`
   - Clock-Abstraktion (`NewClock`, `NewFixedClock`)
   - Unit-Tests für Konvertierung

2. **Woche 3-4: Arithmetik**
   - `Add()`, `Sub()` für alle Einheiten
   - Monatsende-Overflow-Logik
   - Edge-Case-Tests (Schaltjahre, Monatsgrenzen)

3. **Woche 5: Snap-to**
   - `StartOf()`, `EndOf()` für Week, Month, Quarter, Year
   - `Next()`, `Prev()` für Weekdays
   - Tests für Wochenbeginn-Konfiguration

4. **Woche 6: Differenz**
   - `Diff()`, `Duration`-Type
   - `.Days()`, `.Months()`, `.Years()` (int)
   - `.MonthsFloat()`, `.YearsFloat()` (float64)

5. **Woche 7: Human-Format**
   - `.Human()` mit adaptiver Granularität
   - i18n-Infrastruktur (EN, DE)
   - Tests für alle Granularitäts-Stufen

6. **Woche 8: Parsing**
   - `Parse()` (automatisch)
   - `ParseWithLayout()` (explizit)
   - Mehrdeutigkeits-Handling

7. **Woche 9: Formatierung**
   - `Format()` mit Presets (ISO, EU, US, Long, RFC2822)
   - `FormatLayout()` für Custom Layouts
   - `.Lang()` für mehrsprachige Formatierung

8. **Woche 10: Parsing (relativ)**
   - `ParseRelative()` (today, tomorrow, +X days)
   - Tests für alle Ausdrücke

9. **Woche 11: Inspektion**
   - `.WeekNumber()`, `.Quarter()`, `.DayOfYear()`
   - `.IsWeekend()`, `.IsLeapYear()`
   - ISO 8601 Compliance-Tests

10. **Woche 12: Zeitzone & DST**
    - `.In()` für Timezone-Konvertierung
    - DST-Handling-Tests (Add über DST-Grenze)
    - Error-Handling für ungültige Timezones

11. **Woche 13-14: Polishing**
    - Benchmarks optimieren
    - Dokumentation vervollständigen (README, Godoc)
    - Example-Tests schreiben
    - CI/CD-Pipeline finalisieren

**Total: 14 Wochen (ca. 3,5 Monate)**

### Potenzielle Risiken und Herausforderungen

| Risiko | Wahrscheinlichkeit | Impact | Mitigation |
|--------|-------------------|--------|------------|
| **Monatsende-Overflow-Logik fehlerhaft** | Mittel | Hoch | Umfangreiche Edge-Case-Tests, Referenz-Implementierung von Moment.js/Carbon studieren |
| **DST-Handling inkorrekt** | Mittel | Hoch | Tests für alle DST-Umstellungen 2024-2030, Vergleich mit `time.Time` stdlib |
| **Parsing-Ambiguitäten nicht erkannt** | Niedrig | Mittel | Klare Dokumentation, strikte Error-Rückgabe bei Mehrdeutigkeit |
| **Performance-Ziele nicht erreicht** | Niedrig | Mittel | Frühzeitige Benchmarks, Optimierung vor Feature-Freeze |
| **i18n-Daten zu groß (Binary Size)** | Niedrig | Niedrig | Optional: Sub-Package `quando/lang` für lazy-loading |
| **API-Instabilität (Breaking Changes)** | Mittel | Hoch | v0.x.x während Phase 1, Feedback von DatesAPI v2 Team einholen vor v1.0.0 |
| **Zeit-Überziehung durch Scope-Creep** | Mittel | Mittel | Strikte Priorisierung: Must-have vs. Nice-to-have, keine Features außerhalb Spec |

---

## Anhang A: Code-Beispiele

### Beispiel 1: Komplexe Verkettung

```go
// Berechne: "Erster Montag des übernächsten Quartals"
date := quando.Now().
    Add(2, quando.Quarters).
    StartOf(quando.Quarter).
    Next(time.Monday)

fmt.Println(date.Format(quando.ISO))
```

### Beispiel 2: Differenz mit Human-Format

```go
start := quando.MustParse("2025-01-15")
end := quando.Now()

diff := quando.Diff(start.Time(), end.Time())

fmt.Printf("Differenz: %s\n", diff.Human(quando.LangDE))
// Output: "Differenz: 13 Monate, 2 Tage"
```

### Beispiel 3: Zeitzone-Konvertierung mit DST

```go
// UTC-Datum
utcDate := quando.MustParse("2026-03-31T01:00:00Z")

// Konvertierung nach Berlin (CET → CEST-Umstellung am 29. März 2026)
berlinDate, _ := utcDate.In("Europe/Berlin")

// +1 Tag (DST-safe)
nextDay := berlinDate.Add(1, quando.Days)

fmt.Println(nextDay.Format(quando.RFC2822))
// Output: "Wed, 01 Apr 2026 01:00:00 +0200" (CEST)
```

### Beispiel 4: Parsing-Workflow

```go
func ParseUserInput(input string) (quando.Date, error) {
    // Versuche automatisches Parsing
    date, err := quando.Parse(input)
    if err == nil {
        return date, nil
    }

    // Fallback: Relative Ausdrücke
    date, err = quando.ParseRelative(input)
    if err == nil {
        return date, nil
    }

    // Fallback: Explizites Format (EU)
    date, err = quando.ParseWithLayout(input, "02.01.2006")
    if err == nil {
        return date, nil
    }

    return quando.Date{}, fmt.Errorf("unable to parse: %s", input)
}
```

### Beispiel 5: Deterministische Tests

```go
func TestBusinessLogic(t *testing.T) {
    // Fixed Clock für deterministische Tests
    clock := quando.NewFixedClock(time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC))

    // Business-Logik mit injizierter Zeit
    result := CalculateDeadline(clock)

    expected := clock.Now().Add(30, quando.Days).EndOf(quando.Month)
    assert.Equal(t, expected, result)
}

func CalculateDeadline(clock quando.Clock) quando.Date {
    return clock.Now().Add(30, quando.Days).EndOf(quando.Month)
}
```

---

## Anhang B: Vergleich zu time.Time

| Operation | time.Time (stdlib) | quando |
|-----------|-------------------|--------|
| **+2 Monate, -3 Tage** | 5+ Zeilen Code | `quando.Now().Add(2, Months).Sub(3, Days)` |
| **Ende des Quartals** | Manuell rechnen (Quartal bestimmen, letzter Tag) | `quando.Now().EndOf(Quarter)` |
| **Differenz menschenlesbar** | Nicht verfügbar | `Diff(a, b).Human()` |
| **Kalenderwoche** | Nicht verfügbar | `.WeekNumber()` |
| **Quartal** | Nicht verfügbar | `.Quarter()` |
| **Nächster Montag** | Komplexe Loop-Logik | `quando.Now().Next(time.Monday)` |
| **Parsing (automatisch)** | Explizites Layout nötig | `quando.Parse("09.02.2026")` |
| **Mehrsprachige Formatierung** | Nicht verfügbar | `.Lang(DE).Format(Long)` |

**Fazit:** `quando` reduziert typische Datums-Operationen von 5-10 Zeilen auf 1 Zeile, bei gleichbleibender Typ-Sicherheit und Performance.

---

## Anhang C: Migration-Guide (für DatesAPI v2)

### Schritt 1: Dependency hinzufügen

```bash
go get code.beautifulmachines.dev/quando
```

### Schritt 2: Import

```go
import "code.beautifulmachines.dev/quando"
```

### Schritt 3: API-Endpunkte migrieren

**Vorher (DatesAPI v1):**
```go
// Manuelle Datumsberechnung
func HandleAddDays(w http.ResponseWriter, r *http.Request) {
    date := parseDate(r.URL.Query().Get("date"))
    days := parseInt(r.URL.Query().Get("days"))

    result := date.AddDate(0, 0, days)
    writeJSON(w, result)
}
```

**Nachher (DatesAPI v2 mit quando):**
```go
func HandleAddDays(w http.ResponseWriter, r *http.Request) {
    date, _ := quando.Parse(r.URL.Query().Get("date"))
    days := parseInt(r.URL.Query().Get("days"))

    result := date.Add(days, quando.Days)
    writeJSON(w, result)
}
```

**Vorteil:** Konsistente API, weniger Fehleranfälligkeit, bessere Lesbarkeit.

---

## Anhang D: Performance-Benchmarks (Zielwerte)

```
BenchmarkAdd-8                  1000000000   0.8 ns/op   0 B/op   0 allocs/op
BenchmarkSub-8                  1000000000   0.8 ns/op   0 B/op   0 allocs/op
BenchmarkStartOf-8              500000000    1.2 ns/op   0 B/op   0 allocs/op
BenchmarkEndOf-8                500000000    1.2 ns/op   0 B/op   0 allocs/op
BenchmarkDiff-8                 500000000    1.5 ns/op   0 B/op   0 allocs/op
BenchmarkDiffFloat-8            300000000    2.0 ns/op   0 B/op   0 allocs/op
BenchmarkFormat-8               200000000    8.0 ns/op   64 B/op  1 allocs/op
BenchmarkParse-8                100000000    9.5 ns/op   48 B/op  1 allocs/op
BenchmarkParseRelative-8        50000000     18.0 ns/op  96 B/op  2 allocs/op
BenchmarkHuman-8                100000000    12.0 ns/op  128 B/op 2 allocs/op
```

**Ziel:** Alle Operationen unter 20 ns/op, minimale Allocations.

---

## Änderungshistorie

| Version | Datum | Änderungen |
|---------|-------|------------|
| 1.0 | 2026-02-09 | Initiales PRD erstellt |
