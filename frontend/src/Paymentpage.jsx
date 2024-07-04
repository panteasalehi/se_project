import React, { useState } from 'react';



export const Paymentpage = (props) => {
    const id = props.data;
    const[cardnum1 , setCardnum1] = useState('')
    const[cardnum2, setCardnum2] = useState('')
    const[cardnum3 , setCardnum3] = useState('')
    const[cardnum4 , setCardnum4] = useState('')
    const [password, setPassword] = useState('');
    const [cvv2, setCvv2] = useState('');
    const [month, setMonth] = useState('');
    const [year, setYear] = useState('');
    
    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        

        if(cardnum1 != '' && cardnum2 != '' && cardnum3!= '' && cardnum4 != '' && password != '' &&  cvv2 != '' && month != '' && year != ''){
            const regobj = {
                UserID: id,
                Status: "true"
            }
        

            props.onFormSwitch("mainpage",id)
        }
        else{
            alert("Please complete all fields");
        }
        
    }

    
        
    

    return (
        <div style={{textAlign:"center"}} >  
          شماره حساب
          <button  onClick = {handleSubmit} style={{ backgroundColor:"green" ,margin:"1%", padding:"0.75%" , width:"15%"}}>ورود به صفحه پرداخت</button>
          <button  onClick = {() => props.onFormSwitch("register",id)} style={{ backgroundColor:"red",margin:"2%",marginTop:"2%",marginRight:"23%", padding:"0.75%" , width:"15%"}}>انصراف</button>
              
              
        </div> 

        
        
    )
  }
  
  

  

