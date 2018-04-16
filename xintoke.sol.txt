pragma solidity ^0.4.0;

contract xintoken {
    
    struct rate
    {
        uint  v_rate;//1~5
        string comment;
    }
    struct repair 
    {
        string url;
        mapping(address => uint) buy;
        mapping(address => rate) fallow;
    }
    mapping(address => mapping(address => repair)) public info;
    
     function xintoken() public {
    }
    
    function add_repair(address car, string url) public payable {
        info[car][msg.sender].url = url;
    }
    
    function add_rate(address car, address sell, uint irate, string comment) public payable {
        info[car][sell].fallow[msg.sender].v_rate = irate;
        info[car][sell].fallow[msg.sender].comment = comment;
    }

    function buy(address car, address sell) public payable returns (string url){
        info[car][sell].buy[msg.sender]=1;
        return info[car][sell].url;
    }
}