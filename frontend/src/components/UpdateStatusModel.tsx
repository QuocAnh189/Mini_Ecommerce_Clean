//hooks
import { useState } from 'react'

//enums
import { EStatusOrder } from '@constants/enum'

//components
import toast from 'react-hot-toast'
import Loading from './Loading'

//redux
import { useUpdateStatusOrderMutation } from '@redux/services/order'

interface IProps {
  orderId: string
  currentStatus: EStatusOrder
}
const UpdateStatusModel = (props: IProps) => {
  const { orderId, currentStatus } = props
  const [status, setStatus] = useState<EStatusOrder>(currentStatus)
  const [UpdateStatusOrder, { isLoading }] = useUpdateStatusOrderMutation()

  const handleUpdateStatus = async () => {
    try {
      const result = await UpdateStatusOrder({ orderId, status })
      if (result) {
        const modal = document.getElementById(`update_status_modal_${orderId}`) as HTMLDialogElement
        modal.close()
        toast.success('Change status successfully.')
      }
    } catch (error) {
      toast.error('Something went wrong.')
    }
  }

  return (
    <dialog id={`update_status_modal_${orderId}`} className="modal">
      <div className="modal-box">
        <h3 className="text-lg font-bold mb-4">Change Status</h3>
        <select defaultValue={status} onChange={(e) => setStatus(e.target.value as EStatusOrder)} className="select">
          <option value="">Pick a Status</option>
          <option value={EStatusOrder.NEW} selected={EStatusOrder.NEW === currentStatus}>
            New
          </option>
          <option value={EStatusOrder.PROGRESS} selected={EStatusOrder.PROGRESS === currentStatus}>
            Progress
          </option>
          <option value={EStatusOrder.DONE} selected={EStatusOrder.DONE === currentStatus}>
            Done
          </option>
          <option value={EStatusOrder.CANCELED} selected={EStatusOrder.CANCELED === currentStatus}>
            Canceled
          </option>
        </select>
        <div className="modal-action">
          <form method="dialog" className="flex gap-4">
            <button onClick={handleUpdateStatus} className="btn btn-success">
              {isLoading ? <Loading /> : 'Update'}
            </button>
            <button className="btn">Close</button>
          </form>
        </div>
      </div>
    </dialog>
  )
}

export default UpdateStatusModel
