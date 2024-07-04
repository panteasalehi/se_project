import React, { useState } from 'react';
import { toast } from "react-toastify";
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";

export const Mainpage = (props) => {
    const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const email = props.data;
    const { items, loading, error} = state;
    const [store,setstore]= useState("ندارد")
    const [elevator,setelevator]= useState("ندارد")
    const [parking,setparking]= useState("ندارد")
  
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
    const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        fetch("http://localhost:8000/user/" + email).then((res) => {
            return res.json();
        }).then((resp) => {
            //console.log(resp)
            if (Object.keys(resp).length === 0) {
                toast.error('Please Enter valid email');
            } else {
                console.log(resp.seller );
                if (resp.seller === "true") {
                    toast.success('welcome');
                    props.onFormSwitch("makepost",email)
                    
                    //history.push("/mainpage", { email: email })
                    //usenavigate('/')
                }else{
                    toast.error('customers do not have access to this section');
                }
            }
        }).catch((err) => {
            toast.error('process Failed due to :' + err.message);
        });
           
        
    }
    
    return (
      
        <div >  
           <button className = "newpost-button"  onClick={() => props.onFormSwitch("myprofilepage",email)} >پروفایل من</button>
            <button className = "myprofile-button" onClick={handleSubmit}>ثبت اگهی</button>
            <div className = "split left" >
        
                <ul className="flex flex-col">
                    <h2 className="text-3xl my-4">لیست اگهی ها </h2>
                    {
                    items.map((item) => (
                        <li style={{backgroundColor:"#E8DFCA"  , margin : "70px"}} onClick={() => props.onFormSwitch("postdetail",email,item.id , item.title) }
                         
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

            <div>
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
                قیمت
                <br></br>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداقل را وارد کنید"></input>
                <br/>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداکثر را وارد کنید" ></input>
                مساحت
                <br></br>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداقل را وارد کنید"></input>
                <br/>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداکثر را وارد کنید" ></input>
                تعداد اتاق
                
                <br></br>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداقل را وارد کنید"></input>
                <br/>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداکثر را وارد کنید" ></input>
                طبقه
                
                <br></br>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداقل را وارد کنید"></input>
                <br/>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداکثر را وارد کنید" ></input>
                <br/>
                سال ساخت 
                <br></br>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداقل را وارد کنید"></input>
                <br/>
                <input style={{font:"small-caption", margin:"1%", padding:"5%" , width:"80%" }} placeholder="حداکثر را وارد کنید" ></input>

                <br/>
                <input onClick={(e)=>setstore("دارد") } type="radio" value="store"  />انباری
                <br></br>
                <br/>
                <input onClick={(e)=>setelevator("دارد") } type="radio" value="elevator"  />اسانسور
                <br></br>
                <br/>
                <input onClick={(e)=>setparking("دارد") } type="radio" value="parking"  />پارکینگ
                </div>
            
                 
            </div>
        </div> 
        
    )
  }
  
  

  

