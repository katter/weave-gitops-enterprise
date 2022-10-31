import { Alert } from '@material-ui/lab';
import { LoadingPage } from '@weaveworks/weave-gitops';
import {
  CanaryParams,
  useListFlaggerObjects,
} from '../../../../contexts/ProgressiveDelivery';
import { AlertListErrors } from '../../../Layout/AlertListErrors';
import { ManagedObjectsTable } from './ManagedObjectsTable';

type Props = CanaryParams;

const ListManagedObjects = (props: Props) => {
  const { error, data, isLoading } = useListFlaggerObjects(props);

  return (
    <>
      <AlertListErrors errors={data?.errors} />
      {isLoading && <LoadingPage />}
      {error && <Alert severity="error">{error.message}</Alert>}
      {data?.objects && <ManagedObjectsTable objects={data.objects} />}
    </>
  );
};

export default ListManagedObjects;
