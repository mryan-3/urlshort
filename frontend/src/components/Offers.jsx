import React from "react";
import {AiOutlineSafety} from 'react-icons/ai'
import { Link,Share2,CheckCheck, MonitorSmartphone, ThumbsUp } from 'lucide-react'


const Offers = () => {
    return(
    <div className="w-full items-center text-[#e3f6f5] ">
            <h3 className="text-center font-bold font-mono text-4xl py-10  text-[#e3f6f5]">What we offer</h3>
        <div className="flex flex-col md:grid grid-rows-2 grid-cols-3 ">
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#bae8e8] text-[#272643]">
            <ThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl  bg-[#bae8e8] text-[#272643]">
            <AiOutlineSafety size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#bae8e8] text-[#272643]">
            <Link size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#bae8e8] text-[#272643]">
            <MonitorSmartphone size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#bae8e8] text-[#272643]">
            <Share2 size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#bae8e8] text-[#272643]">
            <CheckCheck size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        </div>
    </div>
    )
}

export default Offers
