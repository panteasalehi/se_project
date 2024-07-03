import React, { useState } from 'react';
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";


export const Myprofilepage = (props) => {
    const email = props.data;
    const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const[name , setName] = useState('')

    
    

    

    fetch("http://localhost:8000/user/" + email).then((res) => {
      return res.json();
    }).then((resp) => {
      setName(resp.name);
    })
   

    return (
        <div >  
          <div className = "split right" style = {{ padding : "1%" }} >
          <button onClick = {() => props.onFormSwitch("mainpage",email)} style={{  backgroundColor : 'lightpink'}}>بازگشت</button>
          <br/>
          <strong >welcome {name} :)</strong> 
          <br/>
          
          <br></br>
       

            <button style= {{width : "60%", marginBottom:"5%"}} onClick = {() => props.onFormSwitch("Myposts",email)}  >اگهی ها </button> 
            
            <button style= {{width : "60%", marginBottom:"5%"}} >پرداخت ها </button> 
            
            <button  style= {{width : "60%", marginBottom:"5%"}}  onClick = {() => props.onFormSwitch("Myfavourites",email)} >علاقه مندی ها </button> 
             
            <button onClick = {() => props.onFormSwitch("notes",email)} style= {{width : "60%", marginBottom:"5%"  }}>یادداشت ها </button> 
            
            <button  style= {{width : "60%", marginBottom:"5%" }} onClick = {() => props.onFormSwitch("support",email)}>پشتیبانی </button> 
            
            <button  style= {{width : "60%"}}  onClick = {() => props.onFormSwitch("login",email)}>خروج </button> 
        
            </div>
        </div> 
        
    )
  }
  
  

  

