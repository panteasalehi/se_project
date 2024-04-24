import React from "react"
export const Makepost = () => {
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

        <div className = "map">

        </div>

        <div className = "post-cancle">

        </div>

        </div>

        
        </div>
    )
  }
  
  

  