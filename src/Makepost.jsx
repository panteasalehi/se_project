import React, { useState } from 'react';
import { GoogleMap, useLoadScript, Marker } from '@react-google-maps/api';

const libraries = ['places'];
const mapContainerStyle = {
  width: '50vw',
  height: '50vh',
};
const center = {
  lat: 35.7219, // default latitude
  lng: 51.3347, // default longitude
};


export const Makepost = () => {
    const { isLoaded, loadError } = useLoadScript({
        googleMapsApiKey: 'AIzaSyCNVtpIqQ_AMO-Ov07BInNpgkuinI3Zr6Y',
        libraries,
      });
    
      const [markerPosition, setMarkerPosition] = useState(null);
    
      const handleMapClick = (event) => {
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
      
    return (
        <div class="wrapper">
        <div class="header">logo will be here</div>
        <div class="two">
        <text>انتخاب دسته بندی اگهی</text>
        <hr className="hr"></hr>
        <br></br>
        <br></br>
        <br></br>

            <input type="radio" id="owner" value="owner" name="gender" />فروش مسکونی
            <br></br>
            <br></br>
            <input type="radio" id="owner" value="owner" name="gender" />اجاره مسکونی
            <br></br>
            <br></br>

             <input type="radio" id="owner" value="owner" name="gender" />اجاره تجاری و اداری
             <br></br>
            <br></br>
            <input type="radio" id="owner" value="owner" name="gender" />فروش تجاری و اداری
            <br></br>
            <br></br>
            <input type="radio" id="customer" value="customer" name="gender" />پروژه های ساخت و ساز
        </div>
        <div class="three">

        <div className="conditions">
        <input name="myInput" placeholder="عنوان"/>
        <br></br>
        <br></br>

        <input name="myInput" placeholder="متراژ"/>
        <br></br>
        <br></br>
        <input name="myInput" placeholder="قیمت"/>
        <br></br>
        <br></br>
        <input name="myInput" placeholder="تعداد اتاق"/>
        <br></br>
        <br></br>
        <input name="myInput" placeholder="سال ساخت"/>
        <br></br>
        <br></br>
        <input name="myInput" placeholder="طبقه"/>
        <br></br>
        <br></br>
        <input name="myInput" placeholder="سایر توضیحات"/>
        <br></br>
        <br></br>
        </div>

        </div>
        <div className="four">
            
        <GoogleMap
        mapContainerStyle={mapContainerStyle}
        zoom={10}
        center={center}
        onClick={handleMapClick}
      >
        {markerPosition && <Marker position={markerPosition} />}
      </GoogleMap>
        </div>
        </div>

        
    )
  }
  
  

  