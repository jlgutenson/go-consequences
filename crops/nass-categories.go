package crops

//NASSCropMap is a map of Crops to NASS Crop ID #s
func NASSCropMap() map[string]Crop {
	m := make(map[string]Crop)

	m["1"] = BuildCrop(1, "Corn")
	m["2"] = BuildCrop(2, "Cotton")
	m["3"] = BuildCrop(3, "Rice")
	m["4"] = BuildCrop(4, "Sorghum")
	m["5"] = BuildCrop(5, "Soybeans")
	m["6"] = BuildCrop(6, "Sunflower")
	m["10"] = BuildCrop(10, "Peanuts")
	m["11"] = BuildCrop(11, "Tobacco")
	m["12"] = BuildCrop(12, "Sweet Corn")
	m["13"] = BuildCrop(13, "Pop or Orn Corn")
	m["14"] = BuildCrop(14, "Mint")
	m["21"] = BuildCrop(21, "Barley")
	m["22"] = BuildCrop(22, "Durum Wheat")
	m["23"] = BuildCrop(23, "Spring Wheat")
	m["24"] = BuildCrop(24, "Winter Wheat")
	m["25"] = BuildCrop(25, "Other Small Grains")
	m["26"] = BuildCrop(26, "Dbl Crop WinWht/Soybeans")
	m["27"] = BuildCrop(27, "Rye")
	m["28"] = BuildCrop(28, "Oats")
	m["29"] = BuildCrop(29, "Millet")
	m["30"] = BuildCrop(30, "Speltz")
	m["31"] = BuildCrop(31, "Canola")
	m["32"] = BuildCrop(32, "Flaxseed")
	m["33"] = BuildCrop(33, "Safflower")
	m["34"] = BuildCrop(34, "Rape Seed")
	m["35"] = BuildCrop(35, "Mustard")
	m["36"] = BuildCrop(36, "Alfalfa")
	m["37"] = BuildCrop(37, "Other Hay/Non Alfalfa")
	m["38"] = BuildCrop(38, "Camelina")
	m["39"] = BuildCrop(39, "Buckwheat")
	m["41"] = BuildCrop(41, "Sugarbeets")
	m["42"] = BuildCrop(42, "Dry Beans")
	m["43"] = BuildCrop(43, "Potatoes")
	m["44"] = BuildCrop(44, "Other Crops")
	m["45"] = BuildCrop(45, "Sugarcane")
	m["46"] = BuildCrop(46, "Sweet Potatoes")
	m["47"] = BuildCrop(47, "Misc Vegs & Fruits")
	m["48"] = BuildCrop(48, "Watermelons")
	m["49"] = BuildCrop(49, "Onions")
	m["50"] = BuildCrop(50, "Cucumbers")
	m["51"] = BuildCrop(51, "Chick Peas")
	m["52"] = BuildCrop(52, "Lentils")
	m["53"] = BuildCrop(53, "Peas")
	m["54"] = BuildCrop(54, "Tomatoes")
	m["55"] = BuildCrop(55, "Caneberries")
	m["56"] = BuildCrop(56, "Hops")
	m["57"] = BuildCrop(57, "Herbs")
	m["58"] = BuildCrop(58, "Clover/Wildflowers")
	m["59"] = BuildCrop(59, "Sod/Grass Seed")
	m["60"] = BuildCrop(60, "Switchgrass")
	m["61"] = BuildCrop(61, "Fallow/Idle Cropland")
	m["63"] = BuildCrop(63, "Forest")
	m["64"] = BuildCrop(64, "Shrubland")
	m["65"] = BuildCrop(65, "Barren")
	m["66"] = BuildCrop(66, "Cherries")
	m["67"] = BuildCrop(67, "Peaches")
	m["68"] = BuildCrop(68, "Apples")
	m["69"] = BuildCrop(69, "Grapes")
	m["70"] = BuildCrop(70, "Christmas Trees")
	m["71"] = BuildCrop(71, "Other Tree Crops")
	m["72"] = BuildCrop(72, "Citrus")
	m["74"] = BuildCrop(74, "Pecans")
	m["75"] = BuildCrop(75, "Almonds")
	m["76"] = BuildCrop(76, "Walnuts")
	m["77"] = BuildCrop(77, "Pears")
	m["92"] = BuildCrop(92, "Aquaculture")
	m["152"] = BuildCrop(152, "Shrubland")
	m["204"] = BuildCrop(204, "Pistachios")
	m["205"] = BuildCrop(205, "Triticale")
	m["206"] = BuildCrop(206, "Carrots")
	m["207"] = BuildCrop(207, "Asparagus")
	m["208"] = BuildCrop(208, "Garlic")
	m["209"] = BuildCrop(209, "Cantaloupes")
	m["210"] = BuildCrop(210, "Prunes")
	m["211"] = BuildCrop(211, "Olives")
	m["212"] = BuildCrop(212, "Oranges")
	m["213"] = BuildCrop(213, "Honeydew Melons")
	m["214"] = BuildCrop(214, "Broccoli")
	m["215"] = BuildCrop(215, "Avocados")
	m["216"] = BuildCrop(216, "Peppers")
	m["217"] = BuildCrop(217, "Pomegranates")
	m["218"] = BuildCrop(218, "Nectarines")
	m["219"] = BuildCrop(219, "Greens")
	m["220"] = BuildCrop(220, "Plums")
	m["221"] = BuildCrop(221, "Strawberries")
	m["222"] = BuildCrop(222, "Squash")
	m["223"] = BuildCrop(223, "Apricots")
	m["224"] = BuildCrop(224, "Vetch")
	m["225"] = BuildCrop(225, "Dbl Crop WinWht/Corn")
	m["226"] = BuildCrop(226, "Dbl Crop Oats/Corn")
	m["227"] = BuildCrop(227, "Lettuce")
	m["228"] = BuildCrop(228, "Dbl Crop Triticale/Corn")
	m["229"] = BuildCrop(229, "Pumpkins")
	m["230"] = BuildCrop(230, "Dbl Crop Lettuce/Durum Wheat")
	m["231"] = BuildCrop(231, "Dbl Crop Lettuce/Cantaloupe")
	m["232"] = BuildCrop(232, "Dbl Crop Lettuce/Cotton")
	m["233"] = BuildCrop(233, "Dbl Crop Lettuce/Barley")
	m["234"] = BuildCrop(234, "Dbl Crop Durum Wht/Sorghum")
	m["235"] = BuildCrop(235, "Dbl Crop Barley/Sorghum")
	m["236"] = BuildCrop(236, "Dbl Crop WinWht/Sorghum")
	m["237"] = BuildCrop(237, "Dbl Crop Barley/Corn")
	m["238"] = BuildCrop(238, "Dbl Crop WinWht/Cotton")
	m["239"] = BuildCrop(239, "Dbl Crop Soybeans/Cotton")
	m["240"] = BuildCrop(240, "Dbl Crop Soybeans/Oats")
	m["241"] = BuildCrop(241, "Dbl Crop Corn/Soybeans")
	m["242"] = BuildCrop(242, "Blueberries")
	m["243"] = BuildCrop(243, "Cabbage")
	m["244"] = BuildCrop(244, "Cauliflower")
	m["245"] = BuildCrop(245, "Celery")
	m["246"] = BuildCrop(246, "Radishes")
	m["247"] = BuildCrop(247, "Turnips")
	m["248"] = BuildCrop(248, "Eggplants")
	m["249"] = BuildCrop(249, "Gourds")
	m["250"] = BuildCrop(250, "Cranberries")
	m["254"] = BuildCrop(254, "Dbl Crop Barley/Soybeans")
	return m
}
