pragma solidity ^0.4.0;
import './strings.sol';

contract xintoken {
    using strings for *;
    
    struct rate{
        uint  v_rate;//1~5
        string comment;
    }
    
    struct repair {
        string url;
        uint price;
        uint stamp;
        uint balance;
        address seller;
        
        mapping(address => uint) buy;
        mapping(address => rate) fallow;
    }
    
    mapping(address => mapping(address => repair)) public info;
    mapping(address => uint) public coin;
    address public admin;
    mapping(address => string) public index;
    
    int public error;
    
    function xintoken(address padmin) public {
         admin = padmin;
    }
    
    function addcoin(address buyer, uint value) public payable{
        if(msg.sender != admin)
            return;
        coin[buyer] += value;
    }
    
    function uint2str(uint i) public pure returns (string){
        if (i == 0) return "0";
        uint j = i;
        uint len;
        while (j != 0){
            len++;
            j /= 10;
        }
        string memory ret = new string(len);
        uint retptr;
        uint k = len - 1;
        assembly { 
            retptr := add(ret, 32) 
        }
        retptr += k;
        
        while (i != 0){
            uint x = 48 + i % 10;
            assembly {
                mstore8(retptr, x)
            }
            retptr--;
            i /= 10;
        }
        return ret;
    }
    
    ///添加维保记录
    function add_repair(address car, address fhash, string url, uint price, uint stamp) public payable {
        assert(bytes(info[car][fhash].url).length == 0);
        assert(bytes(url).length != 0);
        info[car][fhash].url = url;
        info[car][fhash].price = price;
        info[car][fhash].seller = msg.sender;
        info[car][fhash].stamp = stamp;
        index[car] = index[car].toSlice().concat(uint2str(stamp).toSlice());
        index[car] = index[car].toSlice().concat(",".toSlice());
        index[car] = index[car].toSlice().concat(uint2str(price).toSlice());
        index[car] = index[car].toSlice().concat(";".toSlice());
        error = 4;
    }
    
    function getaddress(address car) public view returns (string rets){
        return index[car];
    }

    ///添加评分
    function add_rate(address car, address fhash, uint irate, string comment) public payable {
        assert(bytes(info[car][fhash].url).length != 0);
        
        info[car][fhash].fallow[msg.sender].v_rate = irate;
        info[car][fhash].fallow[msg.sender].comment = comment;
    }

    ///购买维保记录
    function buy(address car, address fhash) public payable {
        assert(bytes(info[car][fhash].url).length != 0);
        assert(info[car][fhash].price <= coin[msg.sender]);

            
        info[car][fhash].buy[msg.sender] += info[car][fhash].price;
        info[car][fhash].balance += info[car][fhash].price * 9 /10;
        
        coin[admin] += info[car][fhash].price/10;
        coin[msg.sender] = coin[msg.sender] - info[car][fhash].price;
        error = 5;
    }

    ///4s店检查是否购买过维保记录
    function isbuy(address car, address fhash, address buyer) public view returns (uint){
        if(info[car][fhash].seller != msg.sender)
            return;
        if(info[car][fhash].buy[buyer] / info[car][fhash].price >= 1)
            return 1;
        else
            return 0; 
    }
    
    ///获取url
    function geturl(address car, address fhash) public view returns (string){
        if(info[car][fhash].buy[msg.sender] / info[car][fhash].price < 1)
            return '';
            
        return info[car][fhash].url; 
    }
    
    //提现接口
    function collect() public payable{
        assert(msg.sender == admin);
        admin.transfer(address(this).balance);
    }
    
    function selfcollect(address car, address seller, uint value) public payable returns (uint){
        assert(msg.sender == admin);
        info[car][seller].balance -= value;
        return info[car][seller].balance;
    }
    
    function getcoin(address seller) public view returns (uint){
        if(msg.sender != admin)
            return;
            
        return coin[seller];
    }
    
    function getcoin() public view returns (uint){
        return coin[msg.sender];
    }
}