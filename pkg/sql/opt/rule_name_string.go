// Code generated by "stringer -output=rule_name_string.go -type=RuleName rule_name.go rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

const _RuleName_name = "InvalidRuleNameNumManualRuleNamesEliminateEmptyAndEliminateEmptyOrEliminateSingletonAndOrSimplifyAndSimplifyOrSimplifyFiltersFoldNullAndOrNegateComparisonEliminateNotNegateAndNegateOrCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateScalarGroupByHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryNormalizeAnyFilterNormalizeNotAnyFilterEliminateDistinctEliminateGroupByProjectPushSelectIntoInlinableProjectInlineProjectInProjectEnsureJoinFiltersAndEnsureJoinFiltersPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinSimplifyRightJoinEliminateSemiJoinEliminateAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightEliminateLimitPushLimitIntoProjectPushOffsetIntoProjectEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusEliminateProjectEliminateProjectProjectPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneRowNumberColsCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldNullInEmptyFoldNullNotInEmptyNormalizeInConstFoldInNullEliminateExistsProjectEliminateExistsGroupByEliminateSelectEnsureSelectFiltersAndEnsureSelectFiltersMergeSelectsPushSelectIntoProjectPushSelectIntoJoinLeftPushSelectIntoJoinRightMergeSelectInnerJoinPushSelectIntoGroupByPushLimitIntoScanGenerateIndexScansConstrainScanConstrainLookupJoinIndexScanNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 33, 50, 66, 89, 100, 110, 125, 138, 154, 166, 175, 183, 203, 225, 246, 268, 290, 312, 334, 357, 367, 384, 397, 417, 430, 445, 465, 486, 513, 540, 557, 577, 596, 616, 633, 652, 670, 691, 708, 731, 761, 783, 803, 820, 842, 865, 881, 898, 915, 932, 955, 979, 993, 1013, 1034, 1050, 1062, 1074, 1087, 1098, 1109, 1119, 1130, 1149, 1165, 1188, 1204, 1217, 1232, 1246, 1261, 1278, 1296, 1308, 1324, 1339, 1357, 1367, 1379, 1396, 1412, 1425, 1437, 1450, 1468, 1487, 1505, 1520, 1538, 1554, 1564, 1586, 1608, 1623, 1645, 1664, 1676, 1697, 1719, 1742, 1762, 1783, 1800, 1818, 1831, 1859, 1871}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
