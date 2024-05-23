import React, { useState } from 'react';
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";

export const Mainpage = (props) => {
    const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const email = props.data;
    const { items, loading, error} = state;
  
    console.log(items, loading, error);
  
    useEffect(() => {
      dispatch({type: FETCH_ACTIONS.PROGRESS});
      fetch("http://localhost:8080/mainpage", {
        method: "GET",
        credentials: 'include'
      }).then((res) => {
        if (res.status === 200) {
          dispatch({type: FETCH_ACTIONS.SUCCESS, data: res.data});
        }
        else {
          throw new Error('Failed to post');
        }
      }).catch((err) => {
        dispatch({type: FETCH_ACTIONS.FAILURE});
      });
  
    }, []);



    return (
      
        <div >  
           <button className = "newpost-button"  onClick={() => props.onFormSwitch("myprofilepage",email)} >پروفایل من</button>
            <button className = "myprofile-button" onClick={() => props.onFormSwitch("makepost",email) }>ثبت اگهی</button>
            <div className = "split left" >
        
                <ul className="flex flex-col">
                    <h2 className="text-3xl my-4">لیست اگهی ها </h2>
                    {
                    items.map((item) => (
                        <li style={{backgroundColor:"#E8DFCA"  , margin : "70px"}} onClick={() => props.onFormSwitch("postdetail",email,item.id) }
                         
                        key={item.id}>
                        <p className='my-2 text-xl'>
                            <strong>{item.ads}</strong> 
                        </p>
                        <p className='mb-2 text-lg'>
                            توضیحات: <strong>{item.message}</strong>
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
  
  

  

