package lunar

import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"time"
	"github.com/kataras/iris/core/errors"
)

const (
	MINYEAR int = 1891
	MAXYEAR int = 2100
)

type Lunar struct {
	Year   int
	Month  string
	Date   string
	Nian   string
	EMonth     int
	EDate     int
	Zodiac string
}

/**
  * 将阳历转换为阴历
  * @param year 公历-年
  * @param month 公历-月
  * @param date 公历-日
  */
func SolarToLunar(year int, month int, date int) (Lunar, error) {
	if (year <= MINYEAR || year > MAXYEAR) {
		return Lunar{}, errors.New(fmt.Sprintf("The year dont less than %d and greater than %d!", MINYEAR, MAXYEAR));
	}
	yearData := GetYearData(year - MINYEAR)
	between := GetDaysBetweenSolar(year, month, date, year, int(yearData[1]), int(yearData[2]))
	return GetLunarByBetween(year, between, yearData)
}

/**
  * 将阳历转换为阴历
  * @param mytime 公历时间
  */
func SolarTimeToLunar(mytime time.Time) (Lunar, error) {
	year , month, date:= mytime.Year(), int(mytime.Month()), mytime.Day()
	return SolarToLunar(year, month, date)
}

/**
  * 根据距离正月初一的天数计算阴历日期
  * @param year 阳历年
  * @param between 天数
  */
func GetLunarByBetween(year int, between int, yearDate [4]int64) (Lunar, error) {
	lunar := Lunar{}
	t, e, leapMonth := 0, 0, 0
	m := ""
	if between == 0 {
		lunar.Year = year
		lunar.Month = "正月"
		lunar.Date = "初一"
		lunar.EMonth = 1
		lunar.EDate = 1
	} else {
		if between < 0 {
			year = year - 1
			yearDate = GetYearData(year - MINYEAR)
			between = (GetLunarYearDays(year, yearDate) + between)
		}
		//获取当年每月数天数的数组
		yearMonth := GetLunarYearMonths(year, yearDate)
		monthLen := 12
		//是否闰年
		leapMonth = int(yearDate[0])
		if leapMonth != 0 {
			monthLen = 13
		}
		for i := 0; i < monthLen; i++ {
			if (between == yearMonth[i]) {
				t = i + 2
				e = 1
				break;
			} else if ( between < yearMonth[i]) {
				t = i + 1
				if (i > 0 ) {
					e = between - yearMonth[i-1] + 1
				} else {
					e = between + 1
				}
				break;
			}
		}

		if leapMonth != 0 && t == (leapMonth+1) {
			m = "闰" + GetCapitalMonthNum(t-1)
		} else {
			if ( leapMonth != 0 && leapMonth+1 < t) {
				m = GetCapitalMonthNum(t-1)
			} else {
				m = GetCapitalMonthNum(t)
			}
		}
		lunar.Year = year
		lunar.Month = m
		lunar.Date = GetCapitalDateNum(e)
		lunar.EMonth = t
		lunar.EDate = e
	}
	lunar.Nian = GetLunarYearName(year)
	lunar.Zodiac = GetYearZodiac(year)
	return lunar, nil
}

/**
 * 获取阴历年份数据库
 */
func GetYearData(year int) ([4]int64) {
	datas := [210][4]int64{
		{0, 2, 9, 21936}, {6, 1, 30, 9656}, {0, 2, 17, 9584}, {0, 2, 6, 21168}, {5, 1, 26, 43344}, {0, 2, 13, 59728},
		{0, 2, 2, 27296}, {3, 1, 22, 44368}, {0, 2, 10, 43856}, {8, 1, 30, 19304}, {0, 2, 19, 19168}, {0, 2, 8, 42352},
		{5, 1, 29, 21096}, {0, 2, 16, 53856}, {0, 2, 4, 55632}, {4, 1, 25, 27304}, {0, 2, 13, 22176}, {0, 2, 2, 39632},
		{2, 1, 22, 19176}, {0, 2, 10, 19168}, {6, 1, 30, 42200}, {0, 2, 18, 42192}, {0, 2, 6, 53840}, {5, 1, 26, 54568},
		{0, 2, 14, 46400}, {0, 2, 3, 54944}, {2, 1, 23, 38608}, {0, 2, 11, 38320}, {7, 2, 1, 18872}, {0, 2, 20, 18800},
		{0, 2, 8, 42160}, {5, 1, 28, 45656}, {0, 2, 16, 27216}, {0, 2, 5, 27968}, {4, 1, 24, 44456}, {0, 2, 13, 11104},
		{0, 2, 2, 38256}, {2, 1, 23, 18808}, {0, 2, 10, 18800}, {6, 1, 30, 25776}, {0, 2, 17, 54432}, {0, 2, 6, 59984},
		{5, 1, 26, 27976}, {0, 2, 14, 23248}, {0, 2, 4, 11104}, {3, 1, 24, 37744}, {0, 2, 11, 37600}, {7, 1, 31, 51560},
		{0, 2, 19, 51536}, {0, 2, 8, 54432}, {6, 1, 27, 55888}, {0, 2, 15, 46416}, {0, 2, 5, 22176}, {4, 1, 25, 43736},
		{0, 2, 13, 9680}, {0, 2, 2, 37584}, {2, 1, 22, 51544}, {0, 2, 10, 43344}, {7, 1, 29, 46248}, {0, 2, 17, 27808},
		{0, 2, 6, 46416}, {5, 1, 27, 21928}, {0, 2, 14, 19872}, {0, 2, 3, 42416}, {3, 1, 24, 21176}, {0, 2, 12, 21168},
		{8, 1, 31, 43344}, {0, 2, 18, 59728}, {0, 2, 8, 27296}, {6, 1, 28, 44368}, {0, 2, 15, 43856}, {0, 2, 5, 19296},
		{4, 1, 25, 42352}, {0, 2, 13, 42352}, {0, 2, 2, 21088}, {3, 1, 21, 59696}, {0, 2, 9, 55632}, {7, 1, 30, 23208},
		{0, 2, 17, 22176}, {0, 2, 6, 38608}, {5, 1, 27, 19176}, {0, 2, 15, 19152}, {0, 2, 3, 42192}, {4, 1, 23, 53864},
		{0, 2, 11, 53840}, {8, 1, 31, 54568}, {0, 2, 18, 46400}, {0, 2, 7, 46752}, {6, 1, 28, 38608}, {0, 2, 16, 38320},
		{0, 2, 5, 18864}, {4, 1, 25, 42168}, {0, 2, 13, 42160}, {10, 2, 2, 45656}, {0, 2, 20, 27216}, {0, 2, 9, 27968},
		{6, 1, 29, 44448}, {0, 2, 17, 43872}, {0, 2, 6, 38256}, {5, 1, 27, 18808}, {0, 2, 15, 18800}, {0, 2, 4, 25776},
		{3, 1, 23, 27216}, {0, 2, 10, 59984}, {8, 1, 31, 27432}, {0, 2, 19, 23232}, {0, 2, 7, 43872}, {5, 1, 28, 37736},
		{0, 2, 16, 37600}, {0, 2, 5, 51552}, {4, 1, 24, 54440}, {0, 2, 12, 54432}, {0, 2, 1, 55888}, {2, 1, 22, 23208},
		{0, 2, 9, 22176}, {7, 1, 29, 43736}, {0, 2, 18, 9680}, {0, 2, 7, 37584}, {5, 1, 26, 51544}, {0, 2, 14, 43344},
		{0, 2, 3, 46240}, {4, 1, 23, 46416}, {0, 2, 10, 44368}, {9, 1, 31, 21928}, {0, 2, 19, 19360}, {0, 2, 8, 42416},
		{6, 1, 28, 21176}, {0, 2, 16, 21168}, {0, 2, 5, 43312}, {4, 1, 25, 29864}, {0, 2, 12, 27296}, {0, 2, 1, 44368},
		{2, 1, 22, 19880}, {0, 2, 10, 19296}, {6, 1, 29, 42352}, {0, 2, 17, 42208}, {0, 2, 6, 53856}, {5, 1, 26, 59696},
		{0, 2, 13, 54576}, {0, 2, 3, 23200}, {3, 1, 23, 27472}, {0, 2, 11, 38608}, {11, 1, 31, 19176}, {0, 2, 19, 19152},
		{0, 2, 8, 42192}, {6, 1, 28, 53848}, {0, 2, 15, 53840}, {0, 2, 4, 54560}, {5, 1, 24, 55968}, {0, 2, 12, 46496},
		{0, 2, 1, 22224}, {2, 1, 22, 19160}, {0, 2, 10, 18864}, {7, 1, 30, 42168}, {0, 2, 17, 42160}, {0, 2, 6, 43600},
		{5, 1, 26, 46376}, {0, 2, 14, 27936}, {0, 2, 2, 44448}, {3, 1, 23, 21936}, {0, 2, 11, 37744}, {8, 2, 1, 18808},
		{0, 2, 19, 18800}, {0, 2, 8, 25776}, {6, 1, 28, 27216}, {0, 2, 15, 59984}, {0, 2, 4, 27424}, {4, 1, 24, 43872},
		{0, 2, 12, 43744}, {0, 2, 2, 37600}, {3, 1, 21, 51568}, {0, 2, 9, 51552}, {7, 1, 29, 54440}, {0, 2, 17, 54432},
		{0, 2, 5, 55888}, {5, 1, 26, 23208}, {0, 2, 14, 22176}, {0, 2, 3, 42704}, {4, 1, 23, 21224}, {0, 2, 11, 21200},
		{8, 1, 31, 43352}, {0, 2, 19, 43344}, {0, 2, 7, 46240}, {6, 1, 27, 46416}, {0, 2, 15, 44368}, {0, 2, 5, 21920},
		{4, 1, 24, 42448}, {0, 2, 12, 42416}, {0, 2, 2, 21168}, {3, 1, 22, 43320}, {0, 2, 9, 26928}, {7, 1, 29, 29336},
		{0, 2, 17, 27296}, {0, 2, 6, 44368}, {5, 1, 26, 19880}, {0, 2, 14, 19296}, {0, 2, 3, 42352}, {4, 1, 24, 21104},
		{0, 2, 10, 53856}, {8, 1, 30, 59696}, {0, 2, 18, 54560}, {0, 2, 7, 55968}, {6, 1, 27, 27472}, {0, 2, 15, 22224},
		{0, 2, 5, 19168}, {4, 1, 25, 42216}, {0, 2, 12, 42192}, {0, 2, 1, 53584}, {2, 1, 21, 55592}, {0, 2, 9, 54560},
	}
	return datas[year]
}

/*
 * 判断阳历年是否为闰年
 */
func IsLeapYear(year int) bool {
	return (( year%4 == 0 && year%100 != 0 ) || ( year%400 == 0))
}

/**
 * 获取阴历年的天干地支
 */
func GetLunarYearName(year int) string {
	sky := []string{"庚", "辛", "壬", "癸", "甲", "乙", "丙", "丁", "戊", "己"};
	earth := []string{"申", "酉", "戌", "亥", "子", "丑", "寅", "卯", "辰", "巳", "午", "未"};
	mod10 := int(math.Mod(float64(year), float64(10)))
	mod12 := int(math.Mod(float64(year), float64(12)))
	return fmt.Sprintf("%s%s", sky[mod10], earth[mod12])
}

/**
 * 获取阴历年的生肖
 */
func GetYearZodiac(year int) string {
	zodiac := []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"};
	mod12 := int(math.Mod(float64(year), float64(12)))
	return zodiac[mod12]
}

/**
  * 获取阴历每月的天数的数组
  */
func GetLunarMonths(year int, yearData [4]int64) []int {
	bit := strconv.FormatInt(yearData[3], 2)
	var bit2 string
	var mouthLen int
	if (len(bit) < 16) {
		bit2 = strings.Repeat("0", 16-len(bit)) + bit
	} else {
		bit2 = bit
	}
	bb := strings.Split(bit2, "")
	if yearData[0] == 0 {
		mouthLen = 12
	} else {
		mouthLen = 13
	}
	result := make([]int, mouthLen)
	for i := 0; i < mouthLen; i++ {
		str := string(bb[i])
		bitInt, _ := strconv.Atoi(str)
		result[i] = bitInt + 29
	}
	return result
}

/**
  * 获取农历每年的天数
  */
func GetLunarYearDays(year int, yearDate [4]int64) int {
	monthArray := GetLunarYearMonths(year, yearDate)
	total := 0
	for i := 0; i < len(monthArray); i ++ {
		total = total + monthArray[i]
	}
	return total
}

/**
 * 获取阴历年每个月的年累计天数
 */
func GetLunarYearMonths(year int, yearData [4]int64) []int {
	monthData := GetLunarMonths(year, yearData)
	var mouthLen int
	if yearData[0] == 0 {
		mouthLen = 12
	} else {
		mouthLen = 13
	}
	res := make([]int, mouthLen)
	for i := 0; i < mouthLen; i++ {
		temp := 0
		for j := 0; j <= i; j++ {
			temp = temp + monthData[j];
		}
		res[i] = temp
	}
	return res
}

/**
  * 获取阴历年闰月
  * @param year 年份
  */
func GetLeapMonth(year int) int64 {
	yearData := GetYearData((year - MINYEAR));
	return yearData[0]
}

/**
  * 计算2个阳历日期之间的天数
  * @param year 阳历年
  * @param cmonth
  * @param cdate
  * @param dmonth 阴历正月对应的阳历月份
  * @param ddate 阴历初一对应的阳历天数
  */
func GetDaysBetweenSolar(cyear int, cmonth int, cdate int, dyear int, dmonth int, ddate int) int {
	t1 := time.Date(cyear, time.Month(cmonth), cdate, 1, 0, 0, 0, time.Local)
	t2 := time.Date(dyear, time.Month(dmonth), ddate, 1, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}

/**
  * 获取数字的阴历日的叫法
  * @param num 数字
  * @param isMonth 是否是月份的数字
  */
func GetCapitalDateNum(num int) string {
	dateHash := []string{
		"", "一", "二", "三", "四",
		"五", "六", "七", "八", "九", "十",
	};
	if (num <= 10) {
		return "初" + dateHash[num]
	} else if (num > 10 && num < 20) {
		return "十" + dateHash[num-10];
	} else if (num == 20) {
		return "二十"
	} else if (num > 20 && num < 30 ) {
		return "廿" + dateHash[num-20];
	} else if (num == 30) {
		return "三十"
	} else {
		return ""
	}
}

/**
  * 获取阴历月份名称
  * @param num 数字
  */
func GetCapitalMonthNum(num int) string {
	monthHash := []string{
		"", "正月", "二月", "三月", "四月", "五月", "六月",
		"七月", "八月", "九月", "十月", "冬月", "腊月"}
	if num > 0 && num < len(monthHash) {
		return monthHash[num]
	} else {
		return ""
	}
}


