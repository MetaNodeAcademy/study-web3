// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
在一个有序数组中查找目标值
*/
contract BinarySearch {
    function binarySearch(uint256[] memory arr, uint256 num)
        public
        pure
        returns (bool)
    {
        uint256 left = 0;
        uint256 right = arr.length - 1;
        while (left <= right) {
            uint256 mid = left + (right - left) / 2;
            if (arr[mid] == num) {
                return true;
            }
            if (arr[mid] < num) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return false;
    }
}
