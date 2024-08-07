'use client'
import Messages from '@/components/layout/Messages'
import Rooms from '@/components/layout/Rooms'
import SideBar from '@/components/layout/SideBar'
import { NextPage } from 'next'

const Home: NextPage = () => {
  return (
    <div className='font-sans antialiased h-screen flex'>
      {/* Sidebar / Account Info etc. */}
      <SideBar />
      {/* Room List */}
      <Rooms />
      {/* Chat Content */}
      <Messages />
    </div>
  )
}

export default Home
