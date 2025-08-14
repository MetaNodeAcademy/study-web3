// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    //一个mapping来存储候选人的得票数
    string[] private keys;
    mapping(string => uint256) data;

    //一个vote函数，允许用户投票给某个候选人
    function vote(string calldata candidate) public {
        uint256 votes = data[candidate];
        data[candidate] = votes += 1;
        keys.push(candidate);
    }

    //一个getVotes函数，返回某个候选人的得票数
    function getVotes(string calldata candidate) public view returns (uint256) {
        return data[candidate];
    }

    //一个resetVotes函数，重置所有候选人的得票数
    function resetVotes() public {
        for (uint256 i = 0; i < keys.length; i++) {
            delete data[keys[i]];
        }
        delete keys;
    }
}
