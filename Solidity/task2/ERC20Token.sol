// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
// import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/*
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
合约包含以下标准 ERC20 功能：
    balanceOf：查询账户余额。
    transfer：转账。
    approve 和 transferFrom：授权和代扣转账。
    使用 event 记录转账和授权操作。
    提供 mint 函数，允许合约所有者增发代币。
提示：
    使用 mapping 存储账户余额和授权信息。
    使用 event 定义 Transfer 和 Approval 事件。
    部署到sepolia 测试网，导入到自己的钱包
*/
contract ERC20Token {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    string private _name; //名称
    string private _symbol; //符号
    uint8 private _decimals; //小数位数
    uint256 private _totalSupply; //总发行量
    address private _self;//合约拥有者
    //账户余额
    mapping(address account => uint256) private _balances;
    //授权记录表
    mapping(address account => mapping(address spender => uint256)) private _allowances;
    //XToken XTK
    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
        _decimals = 18;
        _totalSupply = 100000;
        _self = msg.sender;
        // 铸造初始代币给合约部署者（你的 MetaMask 地址）
        mint(_self, _totalSupply * 10 **decimals());
    }

    function name() public view returns (string memory) {
        return _name;
    }

    function symbol() public view returns (string memory) {
        return _symbol;
    }

    function decimals() public view returns (uint8) {
        return _decimals;
    }

    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    //查询账户余额
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    //查询授权额度
    function allowances(address account, address spender) public view returns (uint256) {
        return _allowances[account][spender];
    }

    //转账
    function transfer(address to, uint256 amount) public returns (bool) {
       return  _transfer(msg.sender,to,amount);
    }
    function _transfer(address from, address to, uint256 amount) internal returns (bool) {
        require(to == address(0), "ERC20Token: transfer to the zero address");
        require(from == address(0), "ERC20Token: transfer from the zero address");

        uint256 fromBalance = _balances[from];
        require(fromBalance >= amount, "ERC20Token: transfer amount exceeds balance");
        //执行转账操作
        _balances[from] = fromBalance - amount;
        _balances[to] += amount;
        emit Transfer(from,to,amount);
         return true;
    }

    //授权
    function approve(address spender,uint256 amount) public returns (bool) {
        return _approve(msg.sender,spender,amount);
    }
    function _approve(address owner,address spender,uint256 amount) internal returns (bool) {
        require(owner != address(0), "ERC20Token: approve owner the zero address");
        require(spender != address(0), "ERC20Token: approve spender the zero address");
        require(_balances[owner] >= amount, "ERC20Token: approve amount exceeds balance");
        // 记录授权：owner允许spender花amount的币
        _allowances[owner][spender] = amount;
        emit Approval(owner,spender,amount);
        return true;
    }

    //代扣转账
    function transferFrom(address from,address to,uint256 amount) public returns (bool) {
        bool flag = _approve(from,msg.sender,_allowances[from][msg.sender]-amount);
        if (flag) {
           flag = _transfer(from, to, amount);
        }
        return flag;
    }

    modifier onlyOwner(){
        require(msg.sender == _self, "ERC20Token: only owner allowed");
        _;
    }
    //合约所有者增发代币
    function mint(address account,uint256 amount) internal onlyOwner {
        require(account != address(0), "ERC20Token: mint to the zero address");
        //增加总发行量
        _totalSupply += amount;
        //给目标地址加钱
        _balances[account] += amount;
        //从0地址转过来（代表“新发行”）
        emit Transfer(address(0), account, amount);
    } 

}
