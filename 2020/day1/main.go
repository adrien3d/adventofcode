package main

import (
	"github.com/adrien3d/adventofcode/utils"
)

func day1(input string) {
	testInParsed := utils.ParseIntList(input, "\n")
	for _, first := range testInParsed {
		for _, second := range testInParsed {
			if first+second == 2020 {
				utils.Log(true, "info", "Part 1: Found", first, second, "\tres:", first*second)
			}
			for _, third := range testInParsed {
				if first+second+third == 2020 {
					utils.Log(true, "warn", "Part 2: Found", first, second, third, "\tres:", first*second*third)
				}
			}
		}
	}
}

func main() {
	testInput := "1721\n979\n366\n299\n675\n1456"
	realInput := "1322\n1211\n1427\n1428\n1953\n1220\n1629\n1186\n1354\n1776\n1906\n1849\n1327\n1423\n401\n1806\n1239\n1934\n1256\n1223\n1504\n1365\n1653\n1706\n1465\n1810\n1089\n1447\n1983\n1505\n1763\n1590\n1843\n1534\n1886\n1842\n1878\n1785\n1121\n1857\n1496\n1696\n1863\n1944\n1692\n1255\n1572\n1767\n1509\n1845\n1479\n1935\n1507\n1852\n1193\n1797\n1573\n1317\n1266\n1707\n1819\n925\n1976\n1908\n1571\n1646\n1625\n1719\n1980\n1970\n1566\n1679\n1484\n1818\n1985\n1794\n1699\n1530\n1645\n370\n1658\n1345\n1730\n1340\n1281\n1722\n1623\n1148\n1545\n1728\n1325\n1164\n1462\n1893\n1736\n160\n1543\n1371\n1930\n1162\n2010\n1302\n1967\n1889\n1547\n1335\n1416\n1359\n1622\n1682\n1701\n1939\n1697\n1436\n1367\n1119\n1741\n1466\n1997\n1856\n1824\n1323\n1478\n1963\n1832\n1748\n1260\n1244\n1834\n1990\n1567\n1147\n1588\n1694\n1487\n1151\n1347\n1315\n1502\n546\n730\n1742\n1869\n1277\n1224\n1169\n1708\n1661\n174\n1207\n1801\n1880\n1390\n1747\n1215\n1684\n1498\n1965\n1933\n1693\n1129\n1578\n1189\n1251\n1727\n1440\n1178\n746\n1564\n944\n1822\n1225\n1523\n1575\n1185\n37\n1866\n1766\n1737\n1800\n1633\n1796\n1161\n1932\n1583\n1395\n1288\n1991\n229\n1875\n1540\n1876\n1191\n1858\n1713\n1725\n1955\n1250\n1987\n1724"

	utils.Log(true, "info", "Solving test input")
	day1(testInput)
	utils.Log(true, "info", "Solving real input")
	day1(realInput)
}
