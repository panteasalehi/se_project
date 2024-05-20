import React, { useState } from 'react';
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";
import DefaultImage from "./a.jpg";


export const Postdetail = (props) => {
    const email = props.data;
    const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const { items, loading, error} = state;
    const [avatarURL, setAvatarURL] = useState();
    
  

    console.log(items, loading, error);

  
      useEffect(() => {
        dispatch({type: FETCH_ACTIONS.PROGRESS});
    
        const getItems = async () => {
          try{
            
            let response = await axios.get("http://localhost:8000/posts?id=" + id);
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
  const id = props.data2;
    
    return (
          <div class="container">
        
         <div class="d-flex">
         {
            items.map((item) => (
              <div key={item.id}>
                <div class="split left">
                <button className = "myprofile-button" onClick={() => props.onFormSwitch("mainpage",email)}>بازگشت</button>
                
                
                
                
                <img 
                style = {{width : "40%"}}
                src={item.avatarURL}
                alt ="Avatar"
                className="h-96 w-96 rounded-full" />
                <br/>
                <button style={{margin : '3%' , padding : '8px' , backgroundColor : 'lightskyblue'}}> چت</button>
                <button style={{margin  : '3%' , padding : '8px' , backgroundColor : 'lightpink'}}> افزودن به علاقه مندی ها</button>
                </div>
                <div class="split right" style={{padding:"5%"}}>
                <strong > {item.title}:عنوان اگهی</strong>
                <br/>
                <br/>
                <strong>{item.meterage} :متراژ</strong>
                <br/>
                <br/>
                <strong> {item.price}:قیمت </strong>
                <br/>
                <br/>
                <strong> {item.floor}:طبقه</strong>
                <br/>
                <br/>
                <strong> {item.room}:تعداد اتاق</strong>
                <br/>
                <br/>
                <strong>اسانسور:بعدا</strong>
                <br/>
                <br/>
                <strong>پارکینگ:بعدا </strong>
                <br/>
                <br/>
                <strong>انباری:بعدا</strong>
                <br/>
                <br/>
                <strong> توضیحات:{item.details} </strong>
                        </div>
        
               </div>
             

            ))
         }
        </div>
        </div>
        
    )
  }
  
  

  