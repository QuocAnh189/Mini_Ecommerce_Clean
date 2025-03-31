//interfaces
import { IUser } from '@interfaces/user'

const Profile = () => {
  const user: IUser = JSON.parse(localStorage.getItem('user')!)

  return (
    <div className="h-screen flex flex-col border border-solid border-gray-300 rounded-md items-center">
      <div className="flex flex-col w-2/5 items-center pl-8 gap-4">
        <form className="w-full flex items-center mt-4 gap-12 mt-20">
          <div className="flex flex-col flex-1 gap-8">
            <div className="flex items-center">
              <label className="w-1/5">Name</label>
              <input
                readOnly
                value={user.name}
                className="border border-gray-300 rounded-md p-2"
                placeholder="Enter name"
              />
            </div>

            <div className="flex items-center">
              <label className="w-1/5">Email</label>
              <input
                readOnly
                value={user.email}
                className="border border-gray-300 rounded-md p-2"
                placeholder="Enter email"
              />
            </div>
          </div>
          <div className="flex flex-col items-center justify-center gap-3">
            <div className="avatar">
              <div className="w-24 rounded-full">
                <img src={user.avatar_url} />
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}

export default Profile
