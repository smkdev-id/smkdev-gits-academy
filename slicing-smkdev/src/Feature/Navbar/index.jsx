import smkdev from '../../assets/smkdev.png'
import { Dropdown, Space } from 'antd';
import { items } from '../../constant/constant';
const Navbar = () => {
    return (
        <div>
            <div className="fixed top-0 z-10 w-full">
                <div className="flex mx-auto px-32 py-4 justify-between items-center bg-white shadow-lg">
                    <div className="w-[120px]">
                        <img src={smkdev} alt="smkdev"/>
                    </div>
                    <div>
                        <ul className="flex gap-8 items-center cursor-pointer">
                            {items.map((item) => (
                                <div key={item.id}>
                                    <li className='text-base text-sky-950 font-semibold hover:underline underline-offset-8'>
                                        {item.id === 1 ? (
                                            <Dropdown
                                                menu={{
                                                items: item.label.map((label) => ({
                                                key: label,
                                                label,
                                              })),
                                              selectable: true,
                                            }}
                                            className='text-base text-sky-950 font-semibold hover:underline underline-offset-8'
                                            >
                                            <Space>{item.title}</Space>
                                            </Dropdown>
                                            ) : (
                                            item.title
                                        )}
                                    </li>
                                </div>
                            ))}
                            <li>
                                <button className="bg-sky-950 text-white px-6 py-2 rounded-lg text-sm font-semibold">Dashboard</button>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    )
}   


export default Navbar