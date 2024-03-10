package main

import (
	"fmt"
	"strings"
)

func stage1(player *Player) *Stage {
	stage := Stage{
		strings.ReplaceAll(`{{name}} лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.
У печері темно, тому {{name}} іде стежкою, яка веде від печери в ліс.
У лісі {{name}} натикається на мертве тіло дивної тварини. Він обирає нічого з цим не робити й іти далі.Через деякий час {{name}} приходить до безлюдного табору.
Він вже втомлений і вирішує відпочити, а не йти далі.
У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.
Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.
{{name}} непритомніє. А все могло бути зовсім інакше.`, "{{name}}", player.name),
		fmt.Sprintf(`Ледве розплющивши очі, %s відчув холод, який пронизував усе його тіло, 
а повітря було на стільки вологим, що плечі відчували його вагу. Ледве підвівшись, він спробував
дістатися аптечки, адже місце укусу сильно пекло, але відкривши її побачив лише пляшечку спирту та бинт`, player.name),
		nil,
	}

	choise1 := Stage{
		"1. Взяти спирт",
		fmt.Sprintf(`Взявши спирт і виливши левову частку пляшки на руку, %s нічого не відчув. 
Лише легке шипіння. Біль можливо й був, але його перебивав лютий холод. 
Зрозумівши, що медик з нього ніякий, він сів на розкладений стілець, 
зітхнув і спробував зібратися з думками.`, player.name),
		nil,
	}

	choise2 := Stage{
		"2. Взяти бинт",
		fmt.Sprintf(`Обмотавши місце укусу бинтом, тепло почало наповнювати його тіло, 
але %s розумів, що то ілюзія і відправився на пошуки чогось, що дійсно допоможе з раною`, player.name),
		nil,
	}

	player.stage = &stage
	stage.decisions = []*Stage{&choise1, &choise2}
	return &stage
}

func stage2(player *Player, stage1 *Stage, choice int) *Stage {
	var stage2 *Stage = nil

	player.updateState()

	switch choice {
	case 1: //спирт
		stage2 = stage1.decisions[0]
		player.isBleeding = true
		player.extraStrength += 50
		player.hp -= 50
	case 2: //бинт
		stage2 = stage1.decisions[1]
		player.extraStrength += 20
		player.hp -= 20
	}
	player.warmth -= 30

	stage3_1 := Stage{
		"1. Пошукати дрова",
		strings.ReplaceAll(`{{name}} відчув тепло, але не знайшов звідки воно, аж раптом зрозумів,
що його джерелом є його власна кров з рани. Відсутність медичних навичок давалася взнаки, тому
{{name}} міцно притиснув рукою місце укусу.
Так пройшли найближчі чверть години, коли в під брезентом знайшлися сухі дрова. Залишилося їх розпалити.
Наколовши знайдені дрова, {{name}} використав усі наявні сірники, але зігрівся і відчув легкий прилив сили.`, "{{name}}", player.name),
		nil,
	}

	stage3_2 := Stage{
		"2. Пошукати їжу",
		strings.ReplaceAll(`Переривши абсолютно кожен ящик і намет, {{name}} знайшов лише одну консерву.
Нею виявилася ненависна квасоля в томатному соусі, але звуки з живота змушували його смакові рецептори
зупинити свою роботу на найближчі три хвилини. Рівно стільки знадобилося, щоб проковтнути знайдене.`, "{{name}}", player.name),
		nil,
	}

	stage3_3 := Stage{
		"3. Піти на пошуки резидентів лагеря",
		strings.ReplaceAll(`{{name}} знав, що якщо найближчу годину
не знайде допомогу, то може ніколи не побачити близьких. Аж раптом він побачив людські сліди.
Тут недавно були люди, адже на мокрій землі можна було чітко розгледіти протектор взуття і надпис Lowa. 
Значить я йду у правильному напрямку, подумав {{name}}`, "{{name}}", player.name),
		nil,
	}
	stage2.decisions = []*Stage{&stage3_1, &stage3_2, &stage3_3}

	return stage2
}

func stage3(player *Player, stage2 *Stage, choice int) *Stage {
	var stage3 *Stage = nil

	player.updateState()

	switch choice {
	case 1: //дрова
		player.extraStrength += 10
		player.warmth += 10
		player.hp -= 20
		stage3 = stage2.decisions[0]
	case 2: //їжа
		player.warmth += 10
		player.extraStrength -= 10
		player.isBleeding = false
		stage3 = stage2.decisions[1]
	case 3: //люди
		player.extraStrength += 10
		player.hp -= 10
		stage3 = stage2.decisions[2]
	}

	stage4_1 := Stage{
		"1. Покричати в надії, що резиденти легеря поруч",
		strings.ReplaceAll(`{{name}} відчайдушно крикнув у темряву лісу, але у відповідь почув лише ехо власного голосу.`, "{{name}}", player.name),
		nil,
	}

	stage4_2 := Stage{
		"2. Спробувати поспати",
		strings.ReplaceAll(`Зрозумівши, що лагер ніхто не кине посеред лісу, {{name}} улігся спати.`, "{{name}}", player.name),
		nil,
	}

	stage3.decisions = []*Stage{&stage4_1, &stage4_2}

	return stage3
}

func stage4(player *Player, stage3 *Stage, choice int) *Stage {
	var stage4 *Stage = nil
	var final *Stage = nil

	player.updateState()

	switch choice {
	case 1: //кричить
		player.extraStrength += 10
		player.warmth += 10
		player.hp -= 10
		stage4 = stage3.decisions[0]

		final = &Stage{
			"1. Продовжити...",
			fmt.Sprintf("Віддавши усі сили богам, %s кричав так, що навколишні птахи порозліталися по лісу. Він і сам здивувався тому, на скільки голосно він кликав на допомогу. Це та мерехтіння ліхтарика дало плоди,  його знайшли резиденти лагерю. Виявилося, що один з них досвідчений лікар. Він надав потерпілому першу допомогу і зрозумів, що той - непоганий піддослідний для його дисертації, адже має якусь напрочуд дивну будову м'язів.", player.name),
			nil,
		}
	case 2: //спить
		stage4 = stage3.decisions[1]
		final = &Stage{
			"1. Продовжити...",
			strings.ReplaceAll("{{name}} раптово прокинувся від тріскоту гілок. Розплющивши очі, він побачив трьох людей. \"Ну що ж... Вечеря прийшла до нас сама...\". {{name}} відчув, що пахне смаженим, але додаткові сили дозволили йому відштовхнути цікавих гостей, а раптово гострий нюх - вибратися до найближчого міста.", "{{name}}", player.name),
			nil,
		}
	}
	stage4.decisions = []*Stage{final}

	return stage4
}

func final(player *Player, finalStage *Stage) *Stage {
	player.updateState()
	return finalStage.decisions[0]
}
