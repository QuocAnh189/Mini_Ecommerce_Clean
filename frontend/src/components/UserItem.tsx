import { IUser } from '@interfaces/user'

interface IProps {
  user: IUser
}

const UserItem = (props: IProps) => {
  const { user } = props
  return (
    <tr>
      <td>
        <div className="flex items-center gap-3">
          <div className="avatar">
            <div className="mask mask-squircle h-12 w-12">
              <img src={user.avatar_url} alt="Avatar Tailwind CSS Component" />
            </div>
          </div>
          <div>
            <div className="font-bold">{user.name}</div>
          </div>
        </div>
      </td>
      <td>{user.email}</td>
      <td className="flex items-center gap-2">
        {!user.deleted_at ? (
          <>
            <div aria-label="success" className="status status-success"></div>
            Active
          </>
        ) : (
          <>
            <div aria-label="success" className="status status-error"></div>
            UnActive
          </>
        )}
      </td>
    </tr>
  )
}

export default UserItem
