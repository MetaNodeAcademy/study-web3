// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
整数转罗马数字
*/
contract IntToRoman {
    uint256[13] numArr = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
    string[13] symbol = [
        "M",
        "CM",
        "D",
        "CD",
        "C",
        "XC",
        "L",
        "XL",
        "X",
        "IX",
        "V",
        "IV",
        "I"
    ];

    function intToRoman(uint256 num) public view returns (string memory) {
        string memory roman = "";
        for (uint256 i = 0; i < numArr.length; i++) {
            while (num >= numArr[i]) {
                num -= numArr[i];
                roman = string.concat(roman,symbol[i]);
                // roman = string(abi.encodePacked(roman, symbol[i]));
                if (num == 0) {
                    break;
                }
            }
        }
        return roman;
    }
}
