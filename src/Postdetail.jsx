import React, { useState } from 'react';
import { useReducer, useEffect } from "react";
import { inventoryReducer, initialState } from "./reducers/inventoryReducer";
import { FETCH_ACTIONS } from "./actions";
import axios from "axios";
import DefaultImage from "./a.jpg";
import { toast } from "react-toastify";

export const Postdetail = (props) => {
  const[note , setNote] = useState('')
  const email = props.data;
  const postid = props.data2;
  const title = props.data3;
  const [state, dispatch] = useReducer(inventoryReducer, initialState);
    const { items, loading, error} = state;

   
    
    const handleSubmit = (e) => {
      
      e.preventDefault(); // if we dont call page will be reloded and data will be lost
      
      let regobj = { email , postid ,title};
          
          //console.log(regobj);
          fetch("http://localhost:8000/favs", {
              method: "POST",
              headers: { 'content-type': 'application/json' },
              body: JSON.stringify(regobj)
          }).then((res) => {
              toast.success('posted successfully.')
              props.onFormSwitch("mainpage",email)
          }).catch((err) => {
              toast.error('Failed :' + err.message);
          });
  }
  const handleNote = (e) => {
    e.preventDefault(); // if we dont call page will be reloded and data will be lost
    let regobj = { email , postid ,note};
          
          //console.log(regobj);
          fetch("http://localhost:8000/notes", {
              method: "POST",
              headers: { 'content-type': 'application/json' },
              body: JSON.stringify(regobj)
          }).then((res) => {
              toast.success('note added successfully.')
              props.onFormSwitch("mainpage",email)
          }).catch((err) => {
              toast.error('Failed :' + err.message);
          });
}

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
                style = {{width : "30%" , height : "50%"}}
                src={item.avatarURL}
                alt ="Avatar"
                className="h-96 w-96 rounded-full" />
                <br/>
                <br/>
                <br/>
                <br/>
               
                
                <button style={{margin : '1%' , padding : '8px' , backgroundColor : 'lightskyblue'}}> چت</button>
                <button onClick = {handleSubmit} style={{margin  : '1%' , padding : '8px' , backgroundColor : 'lightpink' }} > افزودن به علاقه مندی ها</button>
                <br/>
                <input style={{font:"small-caption"}} value={note} onChange={(e)=>setNote(e.target.value)} type = "text" placeholder="یادداشت را وارد کنید"  />
                <button  onClick = {handleNote} style={{margin:"1%", font:"small-caption"}}>ثبت یادداشت</button>
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
                <strong> {item.numberOfRooms}:تعداد اتاق </strong>
                <br/>
                <br/>
                <strong>اسانسور: {item.elevator}</strong>
                <br/>
                <br/>
                <strong>پارکینگ: {item.parking} </strong>
                <br/>
                <br/>
                <strong>انباری:{item.store}</strong>
                <br/>
                <br/>
                <strong> توضیحات:{item.description} </strong>
                        </div>
        
               </div>
             

            ))
         }
        </div>
        </div>
        
    )
  }
  
  

  