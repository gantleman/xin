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
        uint price;
        mapping(address => uint) buy;
        mapping(address => rate) fallow;
    }
    mapping(address => mapping(address => repair)) public info;
    
    address admin;
    
    function xintoken(address padmin) public {
         admin = padmin;
    }
    
    ///���ά����¼
    function add_repair(address car, string url, uint price) public payable {
        info[car][msg.sender].url = url;
        info[car][msg.sender].price = price;
    }

    ///�������
    function add_rate(address car, address sell, uint irate, string comment) public payable {
        info[car][sell].fallow[msg.sender].v_rate = irate;
        info[car][sell].fallow[msg.sender].comment = comment;
    }

    ///����ά����¼
    function buy(address car, address sell) public payable returns (string url){
        if(bytes(info[car][sell].url).length == 0)
            return '';
            
        info[car][sell].buy[msg.sender] += msg.value;
        if(msg.value != 0)
            sell.transfer(msg.value * 5 / 10);
        
        return info[car][sell].url;
    }

    ///����Ƿ����ά����¼
    function isbuy(address car, address buyer) public view returns (uint){
        
        return info[car][msg.sender].buy[buyer]/info[car][sell].price >= 1 ? 1:0;
    }
    
    //���ֽӿ�
    function collect() public payable{
        if(msg.sender != admin)
        return;
        
        admin.transfer(address(this).balance);
    }
}