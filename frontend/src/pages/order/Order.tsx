//hook
import { useNavigate } from 'react-router-dom'

//util
import formatDate from '@utils/formatDate'

//components
import Pagination from '@components/Pagination'

const Order = () => {
  const navigate = useNavigate()
  return (
    <div className="h-screen flex flex-col items-center mt-20 gap-4">
      <div className="overflow-x-auto px-40">
        <table className="table table-zebra">
          <thead>
            <tr>
              <th>ID</th>
              <th>Code</th>
              <th>Total Price</th>
              <th>Status</th>
              <th>Created</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <th>1</th>
              <td>Cy Ganderton</td>
              <td>Quality Control Specialist</td>
              <td>Blue</td>
              <td>{formatDate(Date())}</td>
              <th>
                <button onClick={() => navigate('/order/1')} className="btn btn-ghost btn-xs">
                  details
                </button>
              </th>
            </tr>
            <tr>
              <th>2</th>
              <td>Hart Hagerty</td>
              <td>Desktop Support Technician</td>
              <td>Purple</td>
              <td>{formatDate(Date())}</td>
              <th>
                <button className="btn btn-ghost btn-xs">details</button>
              </th>
            </tr>
            <tr>
              <th>3</th>
              <td>Brice Swyre</td>
              <td>Tax Accountant</td>
              <td>Red</td>
              <td>{formatDate(Date())}</td>
              <th>
                <button className="btn btn-ghost btn-xs">details</button>
              </th>
            </tr>
          </tbody>
        </table>
      </div>
      <Pagination />
    </div>
  )
}

export default Order
