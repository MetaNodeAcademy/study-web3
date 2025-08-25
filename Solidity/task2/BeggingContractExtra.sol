// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

/*
额外挑战（可选）
    捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
    捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
    时间限制：添加一个时间限制，只有在特定时间段内才能捐赠。
*/
contract BeggingContract {
    //时间限制（开始时间和结束时间）
    uint256 public _startTime;
    uint256 public _endTime;
    address payable private _owner;
    //所有捐赠者
    mapping(address => uint256) private _donors;
    //所有捐赠者地址
    address[] private _allDonors;
    //捐赠金额前三
    address[3] private _top3Donors;

    event Donation(address indexed donor, uint256 amount);
    event Withdraw(address indexed owner, uint256 amount);
    event FallbackCalled(address sender, uint256 amount, bytes data);
    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );

    receive() external payable {
        donate();
    }

    fallback() external payable {
        require(
            msg.value > 0,
            "BeggingContract: The donation amount must be greater than 0"
        );
        _donors[msg.sender] += msg.value;
        emit FallbackCalled(msg.sender, msg.value, msg.data);
    }

    constructor(uint256 startTime,uint256 endTime) {
        _startTime = startTime;
        _endTime = endTime;
        _owner = payable(msg.sender);
        emit OwnershipTransferred(address(0), msg.sender);
    }

    //用户向合约发送以太币，并记录捐赠信息
    function donate() public payable donationPeriod {
        require(
            msg.value > 0,
            "BeggingContract: The donation amount must be greater than 0"
        );
        _addDonor(msg.sender);
        _donors[msg.sender] += msg.value;
        _updateTop3Donors();
        emit Donation(msg.sender, msg.value);
    }

    //添加新捐赠者到列表
    function _addDonor(address donor) private {
        uint256 amount = _donors[donor];
        if (amount < 1){
            _allDonors.push(donor);
        }
    }

    //更新前三名捐赠者
    function _updateTop3Donors() private {
        _top3Donors = [address(0),address(0),address(0)];
        for (uint256 i = 0;i < _allDonors.length;i++){
            address current = _top3Donors[i];
            uint256 currentAmount = _donors[current];

            for (uint256 j = 0;j < _top3Donors.length; j++){
                if (currentAmount > _donors[_top3Donors[j]]){
                    // 移位后面的捐赠者
                    for (uint k = 2; k > j; k--) {
                        _top3Donors[k] = _top3Donors[k - 1];
                    }
                    _top3Donors[j] = current;
                    break;
                }
               
            }

        }
    }

    // 获取捐赠排行榜前三名
    function getTop3Donors() public view returns (address[3] memory) {
        return _top3Donors;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "BeggingContract: Only owner");
        _;
    }
    modifier donationPeriod(){
        require(
            block.timestamp >= _startTime && block.timestamp <= _endTime, 
            "BeggingContract: Donations are only allowed during the specified period"
        );
        _;
    }

    //合约所有者提取所有资金
    function withdraw() public payable onlyOwner {
        //获取当前合约的总金额
        uint256 amount = address(this).balance;
        require(amount > 0, "BeggingContract: No amount can be withdraw");
        _owner.transfer(amount);
        emit Withdraw(_owner, amount);
    }

    //查询某个地址的捐赠金额
    function getDonation(address donor) public view returns (uint256) {
        return _donors[donor];
    }

    //转移所有者权限
    function transferOwnership(address payable newOwner) public onlyOwner {
        require(
            newOwner != address(0),
            "BeggingContract: New owner is the zero address"
        );
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }

    //查询当前所有者
    function owner() public view returns (address) {
        return _owner;
    }
}
