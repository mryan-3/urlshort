import React from "react";
import {FaThumbsUp} from 'react-icons/fa6'

const Offers = () => {
    return(
    <div className="w-full items-center text-[#e3f6f5] ">
            <h3 className="text-center text-[#e3f6f5]">What we offer</h3>
        <div className="flex flex-col md:grid grid-rows-2 grid-cols-3 ">
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#272643]">
            <FaThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#272643]">
            <FaThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#272643]">
            <FaThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#272643]">
            <FaThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#272643]">
            <FaThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        <div className="flex flex-col items-center border-2 mx-10 md:mx-8 my-3 p-4 rounded-xl bg-[#272643]">
            <FaThumbsUp size={50}/>
            <h3>Easy to use</h3>
            <p>We offer an easy to use and navigate service</p>
        </div>
        </div>
    </div>
    )
}

export default Offers
