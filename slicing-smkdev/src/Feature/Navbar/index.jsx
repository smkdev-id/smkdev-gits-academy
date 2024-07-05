import smkdev from '../../assets/smkdev.png'
// import { DownOutlined } from '@ant-design/icons';
import { Dropdown, Space, Typography } from 'antd';
const Navbar = () => {
    const items = [
        { id:1, key:'1', title:"Learn", label:["Bootcamp", "Expert Class", "Challenges"]}, { id:2, key:'2', title:"Community"}, { id:3, key:'3', title:"Blog"},
    ]
    return (
        <div>
            <div className="py-2 px-4 border-b-2">
                    <div className="flex container mx-auto px-32 flex justify-between items-center">
                    <div className="w-[120px]">
                        <img src={smkdev} alt="smkdev"/>
                    </div>
                    <div>
                        <ul className="flex gap-8 flex items-center cursor-pointer">
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