import React, { useState } from 'react';
import { GoogleMap, useLoadScript, Marker } from '@react-google-maps/api';
import { toast } from "react-toastify";
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";
import { useLocation } from "react-router-dom";

export const Mainpage = (props) => {
    const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const email = props.data;
    const { items, loading, error} = state;
  
    console.log(items, loading, error);
  
    useEffect(() => {
      dispatch({type: FETCH_ACTIONS.PROGRESS});
  
      const getItems = async () => {
        try{
          let response = await axios.get("http://localhost:8000/posts");
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



    return (
      
        <div >  
           <button className = "newpost-button"  onClick={() => props.onFormSwitch("myprofilepage",email)} >پروفایل من</button>
            <button className = "myprofile-button" onClick={() => props.onFormSwitch("makepost",email) }>ثبت اگهی</button>
            <div className = "split left">
        
                <ul className="flex flex-col">
                    <h2 className="text-3xl my-4">لیست اگهی ها </h2>
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
                
            </div>
            <div className = "split right" style = {{ padding : "4%" }} >
            <text>انتخاب دسته بندی اگهی</text>
        <hr className="hr"></hr>
        <br></br>

            <button style= {{width : "80%"}} >فروش مسکونی </button> 
                <br></br>
                <br></br>
            <button style= {{width : "80%"}} >اجاره مسکونی </button> 
                <br></br>
                <br></br>
            <button style= {{width : "80%"}} >اجاره تجاری و اداری </button> 
                <br></br>
                <br></br>
            <button style= {{width : "80%"}} >فروش تجاری و اداری </button> 
                <br></br>
                <br></br>
            <button style= {{width : "80%"}} >پروژه های ساخت و ساز </button> 
                <br></br>
                <br></br>
            
        
            </div>
        </div> 
        
    )
  }
  
  

  

