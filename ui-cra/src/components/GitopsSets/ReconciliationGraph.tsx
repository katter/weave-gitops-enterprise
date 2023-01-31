import { Slider } from '@material-ui/core';
import {
  DirectedGraph,
  Flex,
  Kind,
  RequestStateHandler,
  Spacer,
} from '@weaveworks/weave-gitops';
import * as d3 from 'd3';
import * as React from 'react';
import styled from 'styled-components';
import { GitOpsSet, ObjectRef } from '../../api/gitopssets/types.pb';
import { useGetReconciledTree } from '../../contexts/GitOpsSets';
import { RequestError } from '../../types/custom';

export type Props = {
  className?: string;
  parentObject: GitOpsSet;
  source: ObjectRef;
};

const SliderFlex = styled(Flex)`
  padding-top: ${props => props.theme.spacing.base};
  min-height: 400px;
  min-width: 60px;
  width: 5%;
`;

const PercentFlex = styled(Flex)`
  color: ${props => props.theme.colors.primary10};
  padding: 10px;
  background: rgba(0, 179, 236, 0.1);
  border-radius: 2px;
`;

const GraphDiv = styled.div`
  width: 100%;
  height: 100%;
`;

function ReconciliationGraph({ className, parentObject, source }: Props) {
  //grab data
  const {
    data: objects,
    error,
    isLoading,
  } = parentObject
    ? useGetReconciledTree(
        parentObject.name || '',
        parentObject.namespace || '',
        'GitOpsSet',
        parentObject.inventory,
        parentObject.clusterName,
      )
    : { data: [], error: null, isLoading: false };
  //add extra nodes
  const secondNode = {
    name: parentObject.name,
    namespace: parentObject.namespace,
    suspended: parentObject.suspended,
    conditions: parentObject.conditions,
    type: parentObject.type,
    clusterName: parentObject.clusterName,
    children: objects,
    isCurrentNode: true,
  };

  const rootNode = {
    ...source,
    type: source.kind,
    clusterName: parentObject.clusterName,
    children: [secondNode],
  };
  //graph numbers
  const nodeSize = {
    width: 800,
    height: 300,
    verticalSeparation: 150,
    horizontalSeparation: 100,
  };
  //use d3 to create tree structure
  const root = d3.hierarchy(rootNode, d => d.children);
  const makeTree = d3
    .tree()
    .nodeSize([
      nodeSize.width + nodeSize.horizontalSeparation,
      nodeSize.height + nodeSize.verticalSeparation,
    ])
    .separation(() => 1);
  const tree = makeTree(root);
  const descendants = tree.descendants();
  const links = tree.links();

  //zoom
  const defaultZoomPercent = 85;
  const [zoomPercent, setZoomPercent] = React.useState(defaultZoomPercent);

  //pan
  const [pan, setPan] = React.useState({ x: 0, y: 0 });
  const [isPanning, setIsPanning] = React.useState(false);
  const handleMouseDown = () => {
    setIsPanning(true);
  };
  const handleMouseMove = e => {
    //viewBox change. e.movement is change since previous mouse event
    if (isPanning) setPan({ x: pan.x + e.movementX, y: pan.y + e.movementY });
  };
  const handleMouseUp = () => {
    setIsPanning(false);
  };

  return (
    <RequestStateHandler loading={isLoading} error={error as RequestError}>
      <Flex className={className} wide tall>
        <GraphDiv
          onMouseDown={handleMouseDown}
          onMouseMove={handleMouseMove}
          onMouseUp={handleMouseUp}
          //ends drag event if mouse leaves svg
          onMouseLeave={handleMouseUp}
        >
          <DirectedGraph
            descendants={descendants}
            links={links}
            nodeSize={nodeSize}
            zoomPercent={zoomPercent}
            pan={pan}
          />
        </GraphDiv>
        <SliderFlex tall column align>
          <Slider
            //@ts-ignore
            onChange={(_, value: number) => setZoomPercent(value)}
            defaultValue={defaultZoomPercent}
            orientation="vertical"
            aria-label="zoom"
            min={5}
          />
          <Spacer padding="xs" />
          <PercentFlex>{zoomPercent}%</PercentFlex>
        </SliderFlex>
      </Flex>
    </RequestStateHandler>
  );
}

export default styled(ReconciliationGraph)`
  .MuiSlider-vertical {
    min-height: 400px;
  }
  .MuiSlider-vertical .MuiSlider-track {
    width: 6px;
  }
  .MuiSlider-vertical .MuiSlider-rail {
    width: 6px;
  }
  .MuiSlider-vertical .MuiSlider-thumb {
    margin-left: -9px;
  }
  .MuiSlider-thumb {
    width: 24px;
    height: 24px;
  }
`;
