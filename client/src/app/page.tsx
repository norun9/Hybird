'use client'
import Rooms from '@/components/layout/Rooms'
import SideBar from '@/components/layout/SideBar'
import Messages from '@/components/layout/Messages'
import { NextPage } from 'next'

const Home: NextPage = () => {
  return (
    <>
      <div className='font-sans antialiased h-screen flex'>
        {/* Sidebar / Account Info etc. */}
        <SideBar />
        {/* Room List */}
        <Rooms />
        {/* Chat content */}
        <Messages />
      </div>
    </>
  )
}

export default Home
