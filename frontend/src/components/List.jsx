

const List = () => {
    return (
        <div className='flex justify-center w-full '>
            <div className='relative h-[500px] '>
                <div className='absolute top-0 flex  justify-center'>
                    <div className='left-0 h-[1px] animate-border-width rounded-full bg-gradient-to-r from-[rgba(17,17,17,0)] via-white to-[rgba(17,17,17,0)] transition-all duration-1000' />
                </div>
                <div className='flex h-full items-center w-[350px] md:w-[700px] justify-center rounded-md border border-slate-800 bg-[bae8e8] px-3 py-2'>
                    <p className='text-sm text-slate-200'>Card Content</p>
                </div>
            </div>
        </div>
    )
}

export default List
