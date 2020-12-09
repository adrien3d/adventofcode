package main

import (
	"github.com/adrien3d/adventofcode/2020/utils"
	"strconv"
	"strings"
)

func SumOfNumbers(elts []int, value int) []int {
	for i, _ := range elts {
		sum := 0
		var valsToAdd []int
		for j := i; j < len(elts); j++ {
			if sum == value {
				return valsToAdd
			}
			if sum > value {
				break
			}

			sum += elts[j]
			valsToAdd = append(valsToAdd, elts[j])
		}
	}
	return []int(nil)
}

func solve(input string, v bool) (part1TotalScore, part2TotalScore int) {
	rule := 5
	if len(input) > 100 {
		rule = 25
	}
	lines := strings.Split(input, "\n")
	var data []int
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		data = append(data, value)
	}

	for i := rule; i < len(data); i++ {
		sumOfTwoPrevious := 0
		for j := i - rule; j <= i; j++ {
			for k := i - rule; k <= i; k++ {
				//utils.Log(v, "info", i, j, k, data[i], data[j], data[k], data[j]+data[k])
				if j != k && data[i] == data[j]+data[k] {
					//utils.Log(v, "warn", i, j, k, data[i], data[j], data[k], data[j]+data[k])
					sumOfTwoPrevious = data[i]
				}
			}
		}
		utils.Log(v, "err", data[i], sumOfTwoPrevious)
		if sumOfTwoPrevious == 0 {
			numArr := SumOfNumbers(data, data[i])
			utils.Log(v, "warn", numArr)
			utils.Log(v, "warn", utils.IntArrayMin(numArr), utils.IntArrayMax(numArr), utils.IntArrayMin(numArr)+utils.IntArrayMax(numArr))
			return data[i], utils.IntArrayMin(numArr) + utils.IntArrayMax(numArr)
		}
	}

	return part1TotalScore, part2TotalScore
}

func main() {
	testInput := "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"
	realInput := "10\n21\n30\n41\n7\n31\n43\n13\n18\n16\n8\n27\n48\n49\n19\n46\n50\n37\n9\n26\n42\n15\n23\n47\n14\n17\n20\n29\n21\n92\n22\n31\n76\n24\n25\n28\n30\n73\n41\n33\n32\n34\n45\n38\n35\n36\n37\n70\n39\n50\n42\n51\n43\n46\n47\n49\n61\n52\n58\n53\n63\n102\n65\n66\n67\n68\n87\n71\n88\n73\n75\n76\n95\n85\n140\n89\n90\n115\n96\n100\n101\n118\n128\n116\n131\n129\n132\n246\n208\n135\n139\n165\n164\n148\n161\n151\n166\n179\n181\n205\n289\n186\n196\n277\n201\n217\n268\n290\n247\n316\n261\n328\n312\n286\n449\n287\n327\n299\n309\n317\n403\n347\n367\n377\n382\n387\n397\n413\n615\n749\n529\n508\n533\n534\n547\n843\n573\n664\n616\n586\n596\n608\n714\n676\n1094\n724\n1222\n915\n759\n1193\n784\n1357\n1145\n1459\n1080\n1318\n1142\n1067\n1081\n1120\n1159\n1169\n1182\n1194\n1204\n1272\n1284\n1460\n1400\n1508\n1869\n1543\n1674\n1826\n2454\n1851\n2324\n2222\n2236\n2147\n2262\n2148\n2868\n4087\n2289\n2351\n3656\n3636\n2488\n3226\n2556\n3821\n2908\n3751\n5104\n3217\n3910\n3500\n3677\n3998\n4175\n4369\n7498\n4383\n7407\n6544\n4704\n4640\n5396\n4907\n7136\n5044\n5464\n7128\n5773\n9547\n9217\n6408\n6717\n6894\n7127\n7177\n9219\n8544\n8173\n8879\n9413\n9779\n9023\n13184\n11776\n11881\n10104\n10440\n9951\n10508\n10817\n15301\n12181\n17567\n13302\n13585\n13125\n13611\n14021\n18903\n22478\n16717\n17196\n17052\n18292\n19127\n22132\n18974\n20391\n20459\n24402\n34204\n20948\n27869\n22998\n23942\n25306\n25483\n26427\n32488\n26710\n32738\n27632\n39195\n35009\n44719\n33769\n44280\n43376\n39586\n38101\n39922\n40850\n41339\n41407\n43946\n44890\n46431\n46940\n48304\n49248\n79678\n51910\n66827\n54342\n92250\n107677\n82991\n123841\n126367\n77715\n73355\n126109\n86297\n77687\n89643\n80772\n82189\n121017\n96188\n88836\n93194\n203824\n95244\n97552\n118737\n106252\n121169\n137333\n127697\n154127\n151042\n151070\n155402\n155544\n158459\n159652\n159876\n310946\n162961\n169608\n171025\n175383\n184080\n255840\n186388\n248622\n192796\n201496\n203804\n246434\n299048\n609994\n265030\n278739\n318363\n326425\n306472\n313861\n438615\n343732\n319528\n322837\n332569\n346408\n622471\n355105\n421817\n370468\n379184\n405300\n582988\n716876\n447930\n450238\n668251\n543769\n571502\n578891\n585211\n620333\n793970\n629309\n633389\n1033141\n642365\n652097\n655406\n678977\n701513\n776922\n938093\n1084277\n749652\n855538\n853230\n1662450\n1473563\n1068263\n1205544\n1115271\n2206219\n1221256\n1212280\n1297771\n2600543\n1262698\n1532207\n1275754\n1432328\n1294462\n2199548\n2708082\n1380490\n2768025\n1526574\n1602882\n4241588\n1605190\n2865580\n4806762\n2183534\n2273807\n3878997\n2320815\n2336527\n4473355\n2592233\n3886695\n2538452\n2557160\n4859267\n2807961\n6076237\n3654297\n2674952\n4373215\n2907064\n3131764\n3129456\n3208072\n7540992\n4874979\n7682940\n5990824\n4457341\n4504349\n4594622\n4657342\n4928760\n4893687\n6965448\n5095612\n5213404\n7485920\n5232112\n5582016\n5482913\n5804408\n5806716\n7802694\n6036520\n8440184\n6261220\n6337528\n7865414\n10024372\n8961690\n9114683\n15606388\n9051963\n9523382\n9251964\n13333871\n10107091\n9989299\n10309016\n10327724\n10445516\n11287321\n19595091\n24046572\n15006295\n12141936\n11843236\n14063914\n12297740\n12598748\n14126634\n14202942\n16827104\n18013653\n18638065\n18166646\n20317023\n19697480\n19579688\n19241263\n22288752\n20096390\n20434815\n23429257\n35333893\n37493171\n31030046\n32701979\n23985172\n24140976\n30612401\n33305177\n24896488\n26424374\n32216595\n28329576\n32369588\n38448468\n36180299\n36804711\n39793870\n38820951\n38938743\n39676078\n39337653\n31161678\n40531205\n51047216\n63732025\n48126148\n48881660\n49037464\n50409546\n50565350\n51320862\n53226064\n71308331\n64690358\n62604673\n71037546\n59491254\n142345877\n67341977\n67966389\n87214257\n69982629\n70837756\n70100421\n115468125\n70499331\n140481960\n91096555\n114141571\n97919124\n97007808\n99447010\n99602814\n100974896\n132656747\n104546926\n123326485\n122095927\n124181612\n133642219\n194680943\n264663572\n211319716\n135308366\n137949018\n140083050\n140599752\n167845564\n161196976\n221542937\n300502311\n188104363\n220334293\n196454818\n203993936\n355877919\n199049824\n405049237\n228728538\n226642853\n227873411\n262130630\n345724549\n366725903\n268950585\n273257384\n446977146\n349268734\n305794582\n278548770\n395718975\n308445316\n329042540\n423327606\n384559181\n387154187\n392098299\n400448754\n395504642\n651201017\n425692677\n426923235\n454516264\n455371391\n570575946\n490004041\n531081215\n597993125\n542207969\n547499355\n700543615\n754735217\n627817504\n584343352\n637487856\n755965775\n693004497\n882102340\n1179293381\n771713368\n817790976\n787602941\n854965018\n821197319\n852615912\n880208941\n881439499\n909887655\n945375432\n1427708296\n1131842707\n1240503852\n1393453631\n1089707324\n1636837557\n2264655061\n2518939897\n1212160856\n1221831208\n1330492353\n1448970272\n1944672342\n1971809664\n2093600355\n1697999917\n1670406888\n3166503550\n1673813231\n2372346559\n1732824853\n1761648440\n2035082756\n2150391507\n2077218139\n2221550031\n3365263295\n2301868180\n2311538532\n4236464725\n3796731196\n4554022653\n2552323561\n2661131128\n2892238096\n2779462625\n3119377160\n3344220119\n3371813148\n3368406805\n4738349267\n3494473293\n3406638084\n4046159790\n3767907609\n3810042992\n3838866579\n6938397886\n4227609646\n4298768170\n4523418211\n4613406712\n6005351247\n6862880098\n6576193821\n5213454689\n6491190308\n7184549339\n5440593753\n5671700721\n5898839785\n11314543088\n6712626924\n7599422794\n12496541555\n6901111377\n9668203399\n10622353611\n7577950601\n10415060400\n7648909571\n17198547432\n8526377816\n17019959766\n8822186381\n9136824923\n10054000465\n14577418676\n13477305198\n13476790386\n10885155410\n26156784689\n12341705130\n12153220677\n21300215810\n12611466709\n13613738301\n14312049718\n14479061978\n17663202739\n15427489193\n30675852630\n22568281077\n16400136982\n16175287387\n16471095952\n39039377029\n22140116117\n22299491579\n17959011304\n21290045600\n20939155875\n23038376087\n25197205128\n27090528687\n23226860540\n24494925807\n25955443431\n24764687386\n32871232934\n26225205010\n27925788019\n32438073282\n29906551171\n31602776580\n38743568464\n32575424369\n32646383339\n68945928200\n34134298691\n34430107256\n46904803503\n43238647454\n46054732986\n59682724339\n57843588467\n43977531962\n48424065668\n57832339190\n50720130817\n50989892396\n80896443567\n66669356483\n107121648577\n54150993029\n60363861301\n59528564599\n61509327751\n62481975540\n85753769609\n71318992833\n80484840242\n66780682030\n68564405947\n77668754710\n93958671855\n99413958064\n105583297585\n102575058697\n92401597630\n139904762638\n94697662779\n108822231586\n101710023213\n204285081910\n105140885425\n114514854330\n137988349316\n113679557628\n115660320780\n121873189052\n121037892350\n201522732592\n342273431226\n171139464644\n170274429160\n135345087977\n144449436740\n237533509832\n170070352340\n197984895215\n325322974260\n215735555129\n187099260409\n194111620843\n239147099519\n223337085916\n206850908638\n215389580841\n218820443053\n249024645605\n229339878408\n305415440317\n341413893804\n236698213130\n322444348386\n422586463767\n279794524717\n305619517137\n358238725053\n388890795393\n383596536259\n314519789080\n417448706759\n357169612749\n381210881252\n393950169047\n764807417511\n423451499251\n400962529481\n655933682884\n746060408142\n422240489479\n534959395545\n533340232133\n478364524013\n466038091538\n516492737847\n542317730267\n1168300897621\n585414041854\n594314313797\n715408337802\n772487331652\n671689401829\n1134951203535\n695730670332\n715482318561\n1008355821805\n738380494001\n823203018960\n794912698528\n1262553145989\n824414028732\n1533293192529\n1150053925842\n994857261860\n888278581017\n1161768761870\n1332694807798\n1060352405335\n982530829385\n1058810468114\n1127731772121\n1290044984129\n1179728355651\n1266003715626\n1367420072161\n1387171720390\n1410069895830\n1411212988893\n2132584755227\n1538685337521\n1733237755861\n2113248003089\n1618115717488\n2191834100893\n3294353517097\n2041341297499\n1870809410402\n2188084177456\n1883135842877\n2298348476847\n2537801667951\n3271923093382\n2238538823765\n3500004827388\n2186542240235\n5005160849243\n3307345013125\n2445732071277\n2797241616220\n2754591792551\n2798384709283\n3726769514977\n2949898326414\n3156801055009\n3846485758950\n3351353473349\n3488925127890\n3501251560365\n3753945253279\n3912150707901\n4069678083112\n4057351650637\n4071220020333\n4121674666642\n4425081064000\n4632274311512\n4684270895042\n4941134032786\n7473028139991\n4983783856455\n5704490118965\n5753077084402\n5552976501834\n6106699381423\n7482655604325\n5748283035697\n9595624509313\n6301251799763\n9665227792303\n8546755730642\n6840278601239\n13177518258956\n11781412634025\n7666095961180\n7969502358538\n8701952394624\n10645624151751\n8496301084333\n8753948978154\n10736860940857\n9316545206554\n22518273574882\n9924917889241\n20328168364667\n15148601011147\n11301259537531\n11306053586236\n11659675883257\n11854982417120\n19233162025190\n20535361612179\n26808276894404\n15542230995863\n19241463095795\n14506374562419\n15635598319718\n16162397045513\n16723451336692\n35562225872558\n16671454753162\n17198253478957\n24131899404051\n27490580736838\n18070494184708\n27463656583044\n20617804744085\n21226177426772\n21230971475477\n22607313123767\n23156241954651\n22960935420788\n22965729469493\n23514658300377\n28526437170282\n52041095470659\n37393368520990\n30048605558282\n31229825899111\n55361725303162\n30141972882137\n31797995365231\n32833851798675\n33394906089854\n33869708232119\n51274782985054\n38688298928793\n43833490550539\n39296671611480\n44191906896265\n41843982170857\n44382419381423\n42457148902249\n43838284599244\n46670900255028\n64624731988965\n90729394399452\n66703560030794\n53563263858659\n68736904487075\n60190578440419\n61278431457393\n61371798781248\n61939968247368\n62975824680812\n64011681114256\n124124300489306\n66228757888529\n67264614321973\n77984970540273\n89963081913847\n81140653782337\n134921301295717\n81753820513729\n89128049157277\n110682581369284\n86295433501493\n100234164113687\n115503232106027\n108042699036276\n182767846428000"
	utils.Run(solve, testInput, true)
	utils.Run(solve, realInput, true)
}
