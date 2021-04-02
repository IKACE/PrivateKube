"""
Component loader for reusable components.
"""

import kfp
import os
from kfp.components import func_to_container_op
from kfp.components import InputPath, OutputPath


def run_pipeline():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "subpipeline",
            "run_pipeline.yaml",
        )
    )


def run_subpipeline():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "subpipeline",
            "run_subpipeline.yaml",
        )
    )


def run_subpipeline_serialized():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "subpipeline",
            "run_subpipeline_serialized.yaml",
        )
    )


def download_file_from_uri():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "gs",
            "download_file_from_uri.yaml",
        )
    )


def download_dir_from_uri():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "gs",
            "download_dir_from_uri.yaml",
        )
    )


def upload_file_to_uri():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__), "components_src", "gs", "upload_file_to_uri.yaml"
        )
    )


def release_claim():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "claim",
            "release",
            "release.yaml",
        )
    )


def delete_claim():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "claim",
            "delete",
            "delete.yaml",
        )
    )


def allocate_timespan_claim():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "claim",
            "allocate",
            "allocate-timespan.yaml",
        )
    )


def allocate_user_blocks_claim():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "claim",
            "allocate_user_blocks",
            "allocate-user-blocks.yaml",
        )
    )


def commit_uniform_consume_claim():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "claim",
            "commit",
            "commit-uniform-consume.yaml",
        )
    )


def merge_dicts():
    def merge_dicts_func(
        dict_1_name: str,
        dict_2_name: str,
        dict_1_path: InputPath(),
        dict_2_path: InputPath(),
        dict_output_path: OutputPath(),
    ):
        import yaml
        import os

        with open(dict_1_path) as f1:
            d1 = yaml.load(f1, Loader=yaml.FullLoader)

        with open(dict_2_path) as f2:
            d2 = yaml.load(f2, Loader=yaml.FullLoader)

        pdir = os.path.dirname(dict_output_path)
        if pdir and not os.path.exists(pdir):
            os.makedirs(pdir)

        with open(dict_output_path, "w") as fm:
            dm = {}
            dm[dict_1_name] = d1
            dm[dict_2_name] = d2
            yaml.dump(dm, fm)

    merge_dicts_op = func_to_container_op(
        merge_dicts_func, base_image="privatekube/yaml:latest"
    )

    return merge_dicts_op


def set_dict_key():
    return kfp.components.load_component_from_file(
        os.path.join(
            os.path.dirname(__file__),
            "components_src",
            "artifacts",
            "set_dict_key.yaml",
        )
    )
