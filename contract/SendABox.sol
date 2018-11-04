

pragma solidity ^0.4.21;
//pragma experimental ABIEncoderV2;

import './zeppelin-solidity/contracts/math/SafeMath.sol';

contract SendABox {
    using SafeMath for uint256;
    address public owner;
    address developer;
    address government;
    uint public constant SZABO_PER_WEI = 10000000000000000; //  0^16 개ㅑ 0.01 ether = 1 token 

    uint256 box_idx = 0;
      
    mapping(address => uint256) balances;

    bool private closed;
  
    event ev_SendABoxEvent(uint256 indexed _box_idx , address indexed _sender , uint256 _value , uint256 _token , string _message);

    constructor() public {
        box_idx = 0;
        //require(_token.owner != msg.sender);
        owner = msg.sender;

     }

    modifier onlyOwner() {
        require(owner == msg.sender);
        _;
    }
  
    function Contract_SendABox(string message) public payable {
        
        require(!closed);
   
        uint256 amount = msg.value; //주석처리 했을때 79066 -> 79052  , 58929 -> 58943
        
        box_idx = box_idx + 1;
 
        uint256 tokencnt = amount.div(SZABO_PER_WEI);

        balances[msg.sender] = balances[msg.sender].add(tokencnt);

        emit ev_SendABoxEvent(box_idx , msg.sender , amount, tokencnt , message);
    }

     function balanceOf(address _owner) public view returns (uint256) {
        return balances[_owner];
    }

    function nowBoxid() public view returns (uint256) {
        return box_idx;
    }

}