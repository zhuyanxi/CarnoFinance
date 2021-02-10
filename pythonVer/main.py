import xlrd

CSI_300_GROWTH_INDEX_EXCEL = "csi_300_growth_index.xls"
CSI_300_VALUE_INDEX_EXCEL  = "csi_300_value_index.xls"
print("hello py")

bookGrowth = xlrd.open_workbook(CSI_300_GROWTH_INDEX_EXCEL)
sheetGrowth = bookGrowth.sheet_by_index(0)
growthList = sheetGrowth.col_values(5)[1:]
# print(growthList)

bookValue = xlrd.open_workbook(CSI_300_VALUE_INDEX_EXCEL)
sheetValue = bookValue.sheet_by_index(0)
valueList = sheetValue.col_values(5)[1:]
# print(valueList)


niubiStock = [x for x in growthList if x in valueList]
print(niubiStock, len(niubiStock))
print("--------------------------------")
niubiStock1 = [x for x in growthList if x not in valueList]
print(niubiStock1, len(niubiStock1))
print("--------------------------------")
niubiStock2 = [x for x in valueList if x not in growthList]
print(niubiStock2, len(niubiStock2))




bookValue = xlrd.open_workbook(CSI_300_VALUE_INDEX_EXCEL)


