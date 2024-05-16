import React, { useState,useRef } from 'react';
import { GoogleMap, useLoadScript, Marker } from '@react-google-maps/api';
import { toast } from "react-toastify";
 
import ImageUpload from "./ImageUpload"
const libraries = ['places'];
const mapContainerStyle = {
  margin : "5%",
  width: '90%',
  height: '50vh',
};
const center = {
  lat: 35.7219, // default latitude
  lng: 51.3347, // default longitude
};


export const Makepost = (props) => {
  const fileUploadRef = useRef();
  const email = props.data;
  const[category , setcategory] = useState('')
  const[title , settitle] = useState('')
  const[meterage , setmeterage] = useState('')
  const[price , setprice] = useState('')
  const[room , setroom] = useState('')
  const[year , setyear] = useState('')
  const[floor , setfloor] = useState('')
  const[details , setdetails] = useState('')
  const[long , setlong] = useState('')
  const[lt , setlt] = useState('')
  
    const { isLoaded, loadError } = useLoadScript({
        googleMapsApiKey: 'AIzaSyCNVtpIqQ_AMO-Ov07BInNpgkuinI3Zr6Y',
        libraries,
      });
    
      const [markerPosition, setMarkerPosition] = useState(null);
    
      const handleMapClick = (event) => {
        setlong(event.latLng.lng());
        setlt(event.latLng.lat());
        setMarkerPosition({
          lat: event.latLng.lat(),
          lng: event.latLng.lng(),
          
        });
        
      };
    
      if (loadError) {
        return <div>Error loading maps</div>;
      }
    
      if (!isLoaded) {
        return <div>Loading maps</div>;
      }

      const handleSubmit = (e) => {
        e.preventDefault(); // if we dont call page will be reloded and data will be lost
        let regobj = { category, title ,meterage  , price,room ,year ,floor ,details , lt , long ,email};
            
            //console.log(regobj);
            fetch("http://localhost:8000/posts", {
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
    
    return (
          <div class="container">
         <div class="toppane">logo will be here</div>
         <div class="d-flex">
         <div class="leftpane">
         <GoogleMap
            mapContainerStyle={mapContainerStyle}
            zoom={10}
            center={center}
            onClick={handleMapClick}
          >
            {markerPosition && <Marker position={markerPosition} />}
          </GoogleMap>
              <br></br>
              <button onClick = {handleSubmit}   style={{marginRight : '15px' , padding : '8px' , backgroundColor : 'lightskyblue'}} >ثبت اگهی</button>
              <button onClick = {() => props.onFormSwitch("mainpage",email)} style={{marginLeft : '15px' , padding : '8px' , backgroundColor : 'lightpink'}}>بازگشت</button>
              <div className="flex justify-center">
      <ImageUpload />
    </div>  
              
              
        </div>
        <div class="middlepane">
          <br/>
        <input  onChange={(e)=>settitle(e.target.value)} name="myInput" placeholder="عنوان"/>
        <br></br>
        <br></br>

        <input  onChange={(e)=>setmeterage(e.target.value)} name="myInput" placeholder="متراژ"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setprice(e.target.value)} name="myInput" placeholder="قیمت"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setroom(e.target.value)} name="myInput" placeholder="تعداد اتاق"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setyear(e.target.value)} name="myInput" placeholder="سال ساخت"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setfloor(e.target.value)} name="myInput" placeholder="طبقه"/>
        <br></br>
        <br></br>
        <input  onChange={(e)=>setdetails(e.target.value)} name="myInput" placeholder="سایر توضیحات"/>
        <br></br>
        <br></br>
        </div>
        <div class="rightpane">
        <text>انتخاب دسته بندی اگهی</text>
        <hr className="hr"></hr>
        <br></br>
        <br></br>
        <br></br>

            <input onChange={(e)=>setcategory(e.target.value) } type="radio" value="فروش مسکونی"  />فروش مسکونی
            <br></br>
            <br></br>
            <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="اجاره مسکونی"  />اجاره مسکونی
            <br></br>
            <br></br>

             <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="اجاره تجاری و اداری"  />اجاره تجاری و اداری
             <br></br>
            <br></br>
            <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="فروش تجاری و اداری"   />فروش تجاری و اداری
            <br></br>
            <br></br>
            <input onClick={(e)=>setcategory(e.target.value) } type="radio" value="پروژه های ساخت و ساز"    />پروژه های ساخت و ساز
    
        </div>
      </div>
    </div>
        
    )
  }
  
  

  