//utils
import formatNumber from '@utils/formatNumber'

//components
import Pagination from '@components/Pagination'

const OrderLine = () => {
  return (
    <div className="h-screen px-40 mt-20">
      <div className="overflow-x-auto">
        <table className="table">
          <thead>
            <tr>
              <th>Product</th>
              <th>Description</th>
              <th>Quantity</th>
              <th>Price</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>
                <div className="flex items-center gap-3">
                  <div className="avatar">
                    <div className="mask mask-squircle h-12 w-12">
                      <img
                        src="https://img.daisyui.com/images/profile/demo/2@94.webp"
                        alt="Avatar Tailwind CSS Component"
                      />
                    </div>
                  </div>
                  <div>
                    <div className="font-bold">Hart Hagerty</div>
                  </div>
                </div>
              </td>
              <td>Zemlak, Daniel and Leannon</td>
              <td>10</td>
              <td>{formatNumber(100000)} VND</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div className="flex justify-center mt-10">
        <Pagination />
      </div>
    </div>
  )
}

export default OrderLine
