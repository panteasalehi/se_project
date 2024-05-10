import React, { useState } from 'react';
import { GoogleMap, useLoadScript, Marker } from '@react-google-maps/api';
import { toast } from "react-toastify";
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";


export const Mainpage = (props) => {

    const [state, dispatch] = useReducer(inventoryReducer, initialState);

    const { items, loading, error} = state;
  
    console.log(items, loading, error);
  
    useEffect(() => {
      dispatch({type: FETCH_ACTIONS.PROGRESS});
  
      const getItems = async () => {
        try{
          let response = await axios.get("http://localhost:8000/edibles");
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
// 

    return (
        <div >  
           <button className = "newpost-button">پروفایل من</button>
            <button className = "myprofile-button" onClick={() => props.onFormSwitch("makepost")}>ثبت اگهی</button>
            <div className = "split left">
        
                <ul className="flex flex-col">
                    <h2 className="text-3xl my-4">Item List</h2>
                    {
                    items.map((item) => (
                        <li style={{backgroundColor:"#E8DFCA" }}
                         
                        key={item.id}>
                        <p className='my-2 text-xl'>
                            <strong>{item.name}</strong> {' '} {item.picture} of type <strong>{item.type}</strong>
                            {' '} costs <strong>{item.price}</strong> INR/KG.
                        </p>
                        <p className='mb-2 text-lg'>
                            Available in Stock: <strong>{item.quantity}</strong>
                        </p>

                        </li>
                    ))
                    }
                    
                </ul>
                
            </div>
            <div className = "split right">
            <text>انتخاب دسته بندی اگهی</text>
        <hr className="hr"></hr>
        <br></br>
        <br></br>
        <br></br>

            <input  type="radio" value="فروش مسکونی"  />فروش مسکونی
            <br></br>
            <br></br>
            <input type="radio" value="اجاره مسکونی"  />اجاره مسکونی
            <br></br>
            <br></br>

             <input  type="radio" value="اجاره تجاری و اداری"  />اجاره تجاری و اداری
             <br></br>
            <br></br>
            <input  type="radio" value="فروش تجاری و اداری"   />فروش تجاری و اداری
            <br></br>
            <br></br>
            <input type="radio" value="پروژه های ساخت و ساز"    />پروژه های ساخت و ساز
        
            </div>
        </div> 
        
    )
  }
  
  

  

