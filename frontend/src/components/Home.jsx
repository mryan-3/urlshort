import React from "react";







const Home = () => {
    return (
        <div className="w-full h-[550px] text-[#e3f6f5]">
            <h1 className="text-5xl font-mono font-extrabold text-center py-10">GoTiny</h1>
            <div className="flex justify-center ">
                <div className="text-xl w-[750px] font-serif  text-center leading-relaxed flex justify-center">
                    <p>Go Tiny is a tool that empowers you to create short, consise and shareable links. Say goodbye to long links and hello to streamlined and shareable URLs.
                    </p>
                </div>
            </div>
            <div className="flex justify-center mt-5 p-5">
                <div className="flex flex-col p-4 w-[550px] ">
                    <input type="text" placeholder="Enter the URL" className="text-[#272643] h-10 p-4 rounded-md placeholder-black" />
                    <div className="flex justify-center items-center ">
                        <button className='my-3 w-[200px]  transition-background inline-flex h-12 items-center justify-center rounded-md border border-slate-800 bg-gradient-to-r from-slate-100 via-[#c7d2fe] to-[#8678f9] bg-[length:200%_200%] bg-[0%_0%] px-6 font-medium text-black duration-500 hover:bg-[100%_200%] focus:outline-none focus:ring-2 focus:ring-slate-400 focus:ring-offset-2 focus:ring-offset-slate-50'>
                            Click Me
                        </button>
                    </div>
                </div>
            </div>
            <div className="absolute h-full w-full bg-[linear-gradient(to_right,#b1b1b12e_1px,transparent_1px),linear-gradient(to_bottom,#b1b1b12e_1px,transparent_1px)] bg-[size:14px_24px] [mask-image:radial-gradient(ellipse_50%_50%_at_50%_50%,#000_70%,transparent_100%)]">
        card
            </div>



        </div>
    )
}

export default Home
