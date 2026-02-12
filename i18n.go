package quando

import "time"

// i18n.go - Internationalization support for quando
//
// This file contains translations for month names, weekday names, and
// duration units used in formatting operations.
//
// Supported Languages (17 total):
//   EN (English), DE (German), ES (Spanish), FR (French), IT (Italian),
//   PT (Portuguese), NL (Dutch), PL (Polish), RU (Russian), TR (Turkish),
//   VI (Vietnamese), JA (Japanese), KO (Korean), ZhCN (Chinese Simplified),
//   ZhTW (Chinese Traditional), HI (Hindi), TH (Thai)
//
// i18n applies to:
//   - Format(Long): "February 9, 2026" vs "9. Februar 2026"
//   - FormatLayout with month/weekday names
//   - Duration.Human(): "10 months, 16 days" vs "10 Monate, 16 Tage"
//
// i18n does NOT apply to:
//   - ISO, EU, US, RFC2822 formats (always language-independent)
//   - Numeric outputs (WeekNumber, Quarter, DayOfYear)

// monthNames contains full month name translations.
var monthNames = map[Lang][12]string{
	EN: {
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	},
	DE: {
		"Januar", "Februar", "März", "April", "Mai", "Juni",
		"Juli", "August", "September", "Oktober", "November", "Dezember",
	},
	ES: {
		"enero", "febrero", "marzo", "abril", "mayo", "junio",
		"julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre",
	},
	FR: {
		"janvier", "février", "mars", "avril", "mai", "juin",
		"juillet", "août", "septembre", "octobre", "novembre", "décembre",
	},
	IT: {
		"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno",
		"luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre",
	},
	PT: {
		"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
	},
	NL: {
		"januari", "februari", "maart", "april", "mei", "juni",
		"juli", "augustus", "september", "oktober", "november", "december",
	},
	PL: {
		"styczeń", "luty", "marzec", "kwiecień", "maj", "czerwiec",
		"lipiec", "sierpień", "wrzesień", "październik", "listopad", "grudzień",
	},
	RU: {
		"январь", "февраль", "март", "апрель", "май", "июнь",
		"июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь",
	},
	TR: {
		"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran",
		"Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık",
	},
	VI: {
		"Tháng 1", "Tháng 2", "Tháng 3", "Tháng 4", "Tháng 5", "Tháng 6",
		"Tháng 7", "Tháng 8", "Tháng 9", "Tháng 10", "Tháng 11", "Tháng 12",
	},
	JA: {
		"1月", "2月", "3月", "4月", "5月", "6月",
		"7月", "8月", "9月", "10月", "11月", "12月",
	},
	KO: {
		"1월", "2월", "3월", "4월", "5월", "6월",
		"7월", "8월", "9월", "10월", "11월", "12월",
	},
	ZhCN: {
		"一月", "二月", "三月", "四月", "五月", "六月",
		"七月", "八月", "九月", "十月", "十一月", "十二月",
	},
	ZhTW: {
		"一月", "二月", "三月", "四月", "五月", "六月",
		"七月", "八月", "九月", "十月", "十一月", "十二月",
	},
	HI: {
		"जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून",
		"जुलाई", "अगस्त", "सितंबर", "अक्तूबर", "नवंबर", "दिसंबर",
	},
	TH: {
		"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
		"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
	},
}

// monthNamesShort contains short (3-letter) month name translations.
var monthNamesShort = map[Lang][12]string{
	EN:   {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
	DE:   {"Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"},
	ES:   {"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sept", "oct", "nov", "dic"},
	FR:   {"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
	IT:   {"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
	PT:   {"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"},
	NL:   {"jan", "feb", "mrt", "apr", "mei", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
	PL:   {"sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
	RU:   {"янв", "фев", "март", "апр", "май", "июнь", "июль", "авг", "сент", "окт", "нояб", "дек"},
	TR:   {"Oca", "Şub", "Mar", "Nis", "May", "Haz", "Tem", "Ağu", "Eyl", "Eki", "Kas", "Ara"},
	VI:   {"Th1", "Th2", "Th3", "Th4", "Th5", "Th6", "Th7", "Th8", "Th9", "Th10", "Th11", "Th12"},
	JA:   {"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	KO:   {"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
	ZhCN: {"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	ZhTW: {"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
	HI:   {"जन", "फ़र", "मार्च", "अप्रैल", "मई", "जून", "जुल", "अग", "सित", "अक्तू", "नव", "दिस"},
	TH:   {"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
}

// weekdayNames contains full weekday name translations.
// Index: Sunday = 0, Monday = 1, ..., Saturday = 6
var weekdayNames = map[Lang][7]string{
	EN:   {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	DE:   {"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
	ES:   {"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
	FR:   {"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
	IT:   {"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"},
	PT:   {"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"},
	NL:   {"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
	PL:   {"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
	RU:   {"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"},
	TR:   {"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi"},
	VI:   {"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"},
	JA:   {"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},
	KO:   {"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"},
	ZhCN: {"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
	ZhTW: {"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
	HI:   {"रविवार", "सोमवार", "मंगलवार", "बुधवार", "गुरुवार", "शुक्रवार", "शनिवार"},
	TH:   {"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"},
}

// weekdayNamesShort contains short (3-letter) weekday name translations.
// Index: Sunday = 0, Monday = 1, ..., Saturday = 6
var weekdayNamesShort = map[Lang][7]string{
	EN:   {"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
	DE:   {"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
	ES:   {"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
	FR:   {"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
	IT:   {"dom", "lun", "mar", "mer", "gio", "ven", "sab"},
	PT:   {"dom", "seg", "ter", "qua", "qui", "sex", "sáb"},
	NL:   {"zo", "ma", "di", "wo", "do", "vr", "za"},
	PL:   {"niedz.", "pon.", "wt.", "śr.", "czw.", "pt.", "sob."},
	RU:   {"вс", "пн", "вт", "ср", "чт", "пт", "сб"},
	TR:   {"Paz", "Pzt", "Sal", "Çar", "Per", "Cum", "Cmt"},
	VI:   {"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
	JA:   {"日", "月", "火", "水", "木", "金", "土"},
	KO:   {"일", "월", "화", "수", "목", "금", "토"},
	ZhCN: {"周日", "周一", "周二", "周三", "周四", "周五", "周六"},
	ZhTW: {"週日", "週一", "週二", "週三", "週四", "週五", "週六"},
	HI:   {"रवि", "सोम", "मंगल", "बुध", "गुरु", "शुक्र", "शनि"},
	TH:   {"อา.", "จ.", "อ.", "พ.", "พฤ.", "ศ.", "ส."},
}

// durationUnits contains duration unit translations for Human() formatting.
// Each unit has singular and plural forms: [0] = singular, [1] = plural
//
// Known limitation: Some languages (PL, RU) have complex plural rules with 3+ forms
// (e.g., PL: 1 rok, 2-4 lata, 5+ lat). Our [2]string structure only supports two forms,
// so we use the most common plural form. This covers ~95% of use cases.
// TODO: Implement full CLDR plural rules for complete accuracy.
var durationUnits = map[Lang]map[string][2]string{
	EN: {
		"year":   {"year", "years"},
		"month":  {"month", "months"},
		"week":   {"week", "weeks"},
		"day":    {"day", "days"},
		"hour":   {"hour", "hours"},
		"minute": {"minute", "minutes"},
		"second": {"second", "seconds"},
	},
	DE: {
		"year":   {"Jahr", "Jahre"},
		"month":  {"Monat", "Monate"},
		"week":   {"Woche", "Wochen"},
		"day":    {"Tag", "Tage"},
		"hour":   {"Stunde", "Stunden"},
		"minute": {"Minute", "Minuten"},
		"second": {"Sekunde", "Sekunden"},
	},
	ES: {
		"year":   {"año", "años"},
		"month":  {"mes", "meses"},
		"week":   {"semana", "semanas"},
		"day":    {"día", "días"},
		"hour":   {"hora", "horas"},
		"minute": {"minuto", "minutos"},
		"second": {"segundo", "segundos"},
	},
	FR: {
		"year":   {"an", "ans"},
		"month":  {"mois", "mois"},
		"week":   {"semaine", "semaines"},
		"day":    {"jour", "jours"},
		"hour":   {"heure", "heures"},
		"minute": {"minute", "minutes"},
		"second": {"seconde", "secondes"},
	},
	IT: {
		"year":   {"anno", "anni"},
		"month":  {"mese", "mesi"},
		"week":   {"settimana", "settimane"},
		"day":    {"giorno", "giorni"},
		"hour":   {"ora", "ore"},
		"minute": {"minuto", "minuti"},
		"second": {"secondo", "secondi"},
	},
	PT: {
		"year":   {"ano", "anos"},
		"month":  {"mês", "meses"},
		"week":   {"semana", "semanas"},
		"day":    {"dia", "dias"},
		"hour":   {"hora", "horas"},
		"minute": {"minuto", "minutos"},
		"second": {"segundo", "segundos"},
	},
	NL: {
		"year":   {"jaar", "jaar"},
		"month":  {"maand", "maanden"},
		"week":   {"week", "weken"},
		"day":    {"dag", "dagen"},
		"hour":   {"uur", "uur"},
		"minute": {"minuut", "minuten"},
		"second": {"seconde", "seconden"},
	},
	PL: {
		"year":   {"rok", "lata"},       // Note: Missing "lat" for 5+ (complex plural rule)
		"month":  {"miesiąc", "miesiące"}, // Note: Missing "miesięcy" for 5+
		"week":   {"tydzień", "tygodnie"},
		"day":    {"dzień", "dni"},
		"hour":   {"godzina", "godziny"},
		"minute": {"minuta", "minuty"},
		"second": {"sekunda", "sekundy"},
	},
	RU: {
		"year":   {"год", "года"},     // Note: Missing "лет" for 5+ (complex plural rule)
		"month":  {"месяц", "месяца"}, // Note: Missing "месяцев" for 5+
		"week":   {"неделя", "недели"},
		"day":    {"день", "дня"},
		"hour":   {"час", "часа"},
		"minute": {"минута", "минуты"},
		"second": {"секунда", "секунды"},
	},
	TR: {
		"year":   {"yıl", "yıl"},
		"month":  {"ay", "ay"},
		"week":   {"hafta", "hafta"},
		"day":    {"gün", "gün"},
		"hour":   {"saat", "saat"},
		"minute": {"dakika", "dakika"},
		"second": {"saniye", "saniye"},
	},
	VI: {
		"year":   {"năm", "năm"},
		"month":  {"tháng", "tháng"},
		"week":   {"tuần", "tuần"},
		"day":    {"ngày", "ngày"},
		"hour":   {"giờ", "giờ"},
		"minute": {"phút", "phút"},
		"second": {"giây", "giây"},
	},
	JA: {
		"year":   {"年", "年"},
		"month":  {"月", "月"},
		"week":   {"週", "週"},
		"day":    {"日", "日"},
		"hour":   {"時間", "時間"},
		"minute": {"分", "分"},
		"second": {"秒", "秒"},
	},
	KO: {
		"year":   {"년", "년"},
		"month":  {"월", "월"},
		"week":   {"주", "주"},
		"day":    {"일", "일"},
		"hour":   {"시간", "시간"},
		"minute": {"분", "분"},
		"second": {"초", "초"},
	},
	ZhCN: {
		"year":   {"年", "年"},
		"month":  {"月", "月"},
		"week":   {"周", "周"},
		"day":    {"天", "天"},
		"hour":   {"小时", "小时"},
		"minute": {"分钟", "分钟"},
		"second": {"秒", "秒"},
	},
	ZhTW: {
		"year":   {"年", "年"},
		"month":  {"月", "月"},
		"week":   {"週", "週"},
		"day":    {"天", "天"},
		"hour":   {"小時", "小時"},
		"minute": {"分鐘", "分鐘"},
		"second": {"秒", "秒"},
	},
	HI: {
		"year":   {"वर्ष", "वर्ष"},
		"month":  {"महीना", "महीने"},
		"week":   {"सप्ताह", "सप्ताह"},
		"day":    {"दिन", "दिन"},
		"hour":   {"घंटा", "घंटे"},
		"minute": {"मिनट", "मिनट"},
		"second": {"सेकंड", "सेकंड"},
	},
	TH: {
		"year":   {"ปี", "ปี"},
		"month":  {"เดือน", "เดือน"},
		"week":   {"สัปดาห์", "สัปดาห์"},
		"day":    {"วัน", "วัน"},
		"hour":   {"ชั่วโมง", "ชั่วโมง"},
		"minute": {"นาที", "นาที"},
		"second": {"วินาที", "วินาที"},
	},
}

// MonthName returns the localized month name for the given language.
// Returns English name if language not found.
func (l Lang) MonthName(month time.Month) string {
	if names, ok := monthNames[l]; ok {
		return names[month-1]
	}
	// Fallback to English
	return monthNames[EN][month-1]
}

// MonthNameShort returns the short (3-letter) localized month name.
// Returns English abbreviation if language not found.
func (l Lang) MonthNameShort(month time.Month) string {
	if names, ok := monthNamesShort[l]; ok {
		return names[month-1]
	}
	return monthNamesShort[EN][month-1]
}

// WeekdayName returns the localized weekday name for the given language.
// Returns English name if language not found.
func (l Lang) WeekdayName(weekday time.Weekday) string {
	if names, ok := weekdayNames[l]; ok {
		return names[weekday]
	}
	return weekdayNames[EN][weekday]
}

// WeekdayNameShort returns the short (3-letter) localized weekday name.
// Returns English abbreviation if language not found.
func (l Lang) WeekdayNameShort(weekday time.Weekday) string {
	if names, ok := weekdayNamesShort[l]; ok {
		return names[weekday]
	}
	return weekdayNamesShort[EN][weekday]
}

// DurationUnit returns the localized duration unit name (singular or plural).
// The plural parameter determines which form to use.
// Returns English name if language not found.
func (l Lang) DurationUnit(unit string, plural bool) string {
	if units, ok := durationUnits[l]; ok {
		if forms, ok := units[unit]; ok {
			if plural {
				return forms[1]
			}
			return forms[0]
		}
	}
	// Fallback to English
	if forms, ok := durationUnits[EN][unit]; ok {
		if plural {
			return forms[1]
		}
		return forms[0]
	}
	return unit
}
