//components
import ProtectedLayout from '@layouts/protected'

//interfaces
import { IUser } from '@interfaces/user'

const Profile = () => {
  const user: IUser = JSON.parse(localStorage.getItem('user')!)

  return (
    <ProtectedLayout>
      <div className="card w-full flex justify-center items-center  bg-base-100 shadow-sm">
        <figure className="w-1/5 mt-20">
          <img src={user?.avatar_url} alt="Shoes" />
        </figure>
        <div className="card-body">
          <h2 className="card-title">{user?.name}</h2>
          <p>{user?.email}</p>
        </div>
      </div>
    </ProtectedLayout>
  )
}

export default Profile
