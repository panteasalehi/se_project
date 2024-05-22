import React, { useState } from 'react';
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";


export const Myprofilepage = (props) => {
    const email = props.data;
    const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const[name , setName] = useState('')
    const [show, setShow] = useState(true);
    const { items, loading, error} = state;
    
    
    console.log(items, loading, error);

  
      useEffect(() => {
        dispatch({type: FETCH_ACTIONS.PROGRESS});
    
        const getItems = async () => {
          try{
            
            let response = await axios.get("http://localhost:8080/posts?email=" + email);
            if (response.status === 200) {
              dispatch({type: FETCH_ACTIONS.SUCCESS, data: response.data});
            }
          } catch(err){
              console.error(err);
              dispatch({type: FETCH_ACTIONS.ERROR, error: err.message})
          }
        }
    
        getItems();
    
      }, []);
    
    

    fetch("http://localhost:8000/user/" + email).then((res) => {
      return res.json();
    }).then((resp) => {
      setName(resp.name);
    })
   

    return (
        <div >  
          <div className = "split left">
        
          <button className = "myprofile-button" onClick={() => props.onFormSwitch("mainpage",email)}>بازگشت</button>
          {show === "post"? 
          <ul className="flex flex-col">
          <h2 className="text-3xl my-4">لیست اگهی های من </h2>
          {
          items.map((item) => (
              <li style={{backgroundColor:"#E8DFCA"  , margin : "70px"}}
              
              key={item.id}>
              <p className='my-2 text-xl'>
                  <strong>{item.title}</strong> 
              </p>
              <p className='mb-2 text-lg'>
                  توضیحات: <strong>{item.details}</strong>
              </p>

              </li>
          ))
          }
          </ul>
          :
          show==="Favs"?
        <h1>hi</h1>
        :
        null}
          
                
          </div>
          <div className = "split right" style = {{ padding : "4%" }} >
          <strong >welcome {name} :)</strong> 
          <br/>
          <br/>
          <text>انتخاب دسته بندی اگهی</text>
          <hr className="hr"></hr>
          <br></br>
       

            <button style= {{width : "60%", marginBottom:"5%"}} onClick={() => setShow("post")}  >اگهی ها </button> 
            
            <button style= {{width : "60%", marginBottom:"5%"}}>پرداخت ها </button> 
            
            <button  style= {{width : "60%", marginBottom:"5%"}} onClick={() => setShow("favs")}>علاقه مندی ها </button> 
             
            <button  style= {{width : "60%", marginBottom:"5%"}}>یادداشت ها </button> 
            
            <button  style= {{width : "60%", marginBottom:"5%" }}>پشتیبانی </button> 
            
            <button  style= {{width : "60%"}}>خروج </button> 
        
            </div>
        </div> 
        
    )
  }
  
  

  

