// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
罗马数字转整数
*/
contract RomanToInt {
    function romanToInt(string memory num) public pure returns (uint256) {
        bytes memory bs = bytes(num);
        uint256 len = bs.length;
        uint256 result = 0;

        for (uint256 i = 0; i < len; i++) {
            uint256 pre = getValue(bs[i]);
            if (i + 1 < len) {
                uint256 next = getValue(bs[i + 1]);
                if (pre < next) {
                    result += (next - pre);
                    i++;
                    continue;
                }
            }
            result += pre;
        }
        return result;
    }

    function getValue(bytes1 value) private pure returns (uint256) {
        if (value == "I") return 1;
        if (value == "V") return 5;
        if (value == "X") return 10;
        if (value == "L") return 50;
        if (value == "C") return 100;
        if (value == "D") return 500;
        if (value == "M") return 1000;
        return 0; // 无效字符默认返回0
    }
}
