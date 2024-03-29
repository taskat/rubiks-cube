import * as THREE from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
import Cube from "./cube/cube";
import Piece from "./cube/piece";
import Move from "./cube/move";
import { Side } from "./cube/side";
import { ColorsChecker } from "./cube/color";


const PIECE_MATERIAL = new THREE.MeshPhysicalMaterial({
  vertexColors: true,
  metalness: .5,
  roughness: .1,
  clearcoat: .1,
  reflectivity: .5
})

const BEFORE_DELAY = 500;
const AFTER_DELAY = 5000;

export default class Simulator{
  animationSpeed: number = 500;
  model: THREE.BufferGeometry = new THREE.BufferGeometry();
  modelName: string = "/cube-bevelled.glb";
  cubeSize: number = 3;
  renderer: THREE.WebGLRenderer = new THREE.WebGLRenderer({ antialias: true });
  camera: THREE.PerspectiveCamera = new THREE.PerspectiveCamera(34, 1, 1, 100);
  controls: OrbitControls = new OrbitControls(this.camera, this.renderer.domElement);
  scene: THREE.Scene = new THREE.Scene();
  puzzleGroup: THREE.Group = new THREE.Group();
  animationGroup: THREE.Group = new THREE.Group();
  clock: THREE.Clock = new THREE.Clock();
  animationMixer: THREE.AnimationMixer = new THREE.AnimationMixer(this.animationGroup);
  colorPalette: Map<Side, string[][]> = ColorsChecker;
  cube: Cube = new Cube(this.cubeSize, this.colorPalette);
  moves: Move[] = [];
  constructor() {}

  makeRotationMatrix4(rotationMatrix3: math.Matrix): THREE.Matrix4 {
    let values = new Array(16).fill(0);
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        values[i * 4 + j] = rotationMatrix3.get([i, j]);
      }
    }
    values[15] = 1;
    return new THREE.Matrix4().fromArray(values);
  }

  initRenderer(container: HTMLElement) {
    this.renderer.setPixelRatio(window.devicePixelRatio);
    this.renderer.setSize(container.offsetWidth, container.offsetHeight);
    container.appendChild(this.renderer.domElement);
  }

  addMoves(moves: string[]) {
    this.moves = moves.map((move: string) => this.cube.moves.get(move)) as Move[];
    console.log("moves:", this.moves);
  }

  addResizeListener(w: number, h: number) {
    window.addEventListener("resize", () => {
      this.renderer.setSize(w, h);
      this.camera.aspect = w / h;
      this.camera.updateProjectionMatrix();
    });
  }

  createScene() {
    this.scene.background = new THREE.Color("lightgreen");
    this.scene.add(this.camera);
    this.createLights().forEach(light => this.scene.add(light));
    this.scene.add(this.puzzleGroup);
    this.scene.add(this.animationGroup);
  }

  initCamera(w: number, h: number) {
    this.camera.aspect = w / h;
    this.camera.position.set(3, 3, 12);
    this.camera.lookAt(new THREE.Vector3(0, 0, 0));
    this.camera.updateProjectionMatrix();
  }

  createLights(): THREE.DirectionalLight[] {
    let lights: THREE.DirectionalLight[] = [];
    const LIGHT_COLOR = 0xffffff;
    const LIGHT_INTENSITY = 2;
    const LIGHT_DISTANCE = 4;

    const positions = [
      new THREE.Vector3(0, 0, LIGHT_DISTANCE),
      new THREE.Vector3(0, 0, -LIGHT_DISTANCE),
      new THREE.Vector3(0, LIGHT_DISTANCE, 0),
      new THREE.Vector3(0, -LIGHT_DISTANCE, 0),
      new THREE.Vector3(LIGHT_DISTANCE, 0, 0),
      new THREE.Vector3(-LIGHT_DISTANCE, 0, 0)
    ];

    positions.forEach(position => {
      const light = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
      light.position.copy(position);
      lights.push(light);
    });
    return lights;
  }

  initControls() {
    this.controls.minDistance = 5.0;
    this.controls.maxDistance = 40.0;
    this.controls.enableDamping = true;
    this.controls.dampingFactor = 0.9;
    this.controls.autoRotate = false;
    this.controls.autoRotateSpeed = 1.0;
  }

  async init() {
    await this.load3DModel();
    const container = document.getElementById("visualisation-container") as HTMLElement;
    const w = container.offsetWidth;
    const h = container.offsetHeight;
    this.initRenderer(container);
    this.addResizeListener(w, h);
    this.createScene();
    this.initCamera(w, h);
    this.initControls();
    this.createUiPieces();
    this.animate();
    this.scramble();
  }

  async load3DModel() {
    let p = new Promise((resolve, reject) => {
      const loader = new GLTFLoader();
      loader.load(this.modelName,
        gltf => {
          const bufferGeometry = (gltf.scene.children[0] as THREE.Mesh).geometry;
          resolve(bufferGeometry.toNonIndexed());
        },
        undefined,
        reject)
    });
    let model = await p as THREE.BufferGeometry;
    this.model = model;
  }

  setGeometryVertexColors(piece: Piece): THREE.BufferGeometry {
    const pieceGeoemtry = this.model.clone();
    const normalAttribute = pieceGeoemtry.getAttribute("normal") as THREE.BufferAttribute;

    const colors: number[] = [];

    for (let normalIndex = 0; normalIndex < normalAttribute.count; normalIndex += 3) {

      let arrayIndex = normalIndex * normalAttribute.itemSize;
      const normalX = normalAttribute.array[arrayIndex++];
      const normalY = normalAttribute.array[arrayIndex++];
      const normalZ = normalAttribute.array[arrayIndex++];

      const color = this.cube.getColor(piece, normalX, normalY, normalZ);

      colors.push(color.r, color.g, color.b);
      colors.push(color.r, color.g, color.b);
      colors.push(color.r, color.g, color.b);
    }

    pieceGeoemtry.setAttribute("color", new THREE.Float32BufferAttribute(colors, 3));

    return pieceGeoemtry;
  }

  animate() {
    window.requestAnimationFrame(this.animate.bind(this));
    this.controls.update();
    const delta = this.clock.getDelta() * this.animationMixer.timeScale;
    this.animationMixer.update(delta);
    this.renderer.render(this.scene, this.camera);
  }

  scramble() {
    this.recreateUiPieces();
    this.resetUiPieces(this.cube);
    setTimeout(this.animateMoves.bind(this), BEFORE_DELAY, this.moves);
  }

  resetUiPieces(cube: Cube) {
    cube.pieces.forEach(piece => {
      const uiPiece = this.findUiPiece(piece);
      this.resetUiPiece(uiPiece, piece);
    });
  }

  findUiPiece(piece: Piece): THREE.Object3D {
    return this.puzzleGroup.children.find(child => child.userData.id === piece.id) as THREE.Object3D;
  }

  resetUiPiece(uiPiece: THREE.Object3D, piece: Piece) {
    const isEvenSizedCube = this.cubeSize % 2 === 0;
    const adjustValue = (v: number) => isEvenSizedCube ? v < 0 ? v + 0.5 : v - 0.5 : v;
    uiPiece.position.x = adjustValue(piece.coord.x);
    uiPiece.position.y = adjustValue(piece.coord.y);
    uiPiece.position.z = adjustValue(piece.coord.z);
    uiPiece.setRotationFromMatrix(this.makeRotationMatrix4(piece.accTransform3));
  }

  createUiPiece(piece: Piece): THREE.Object3D {
    const pieceGeometryWithColors = this.setGeometryVertexColors(piece);
    const uiPiece = new THREE.Mesh(pieceGeometryWithColors, PIECE_MATERIAL);
    uiPiece.scale.set(0.5, 0.5, 0.5);
    uiPiece.userData.id = piece.id;
    this.resetUiPiece(uiPiece, piece);
    return uiPiece;
  }

  recreateUiPieces() {
    this.cube = new Cube(this.cubeSize, this.colorPalette);
    this.puzzleGroup.children.forEach(child => {
      let mesh = child as THREE.Mesh;
      mesh.geometry.dispose();
    });
    this.scene.remove(this.puzzleGroup);
    this.scene.remove(this.animationGroup);
    this.puzzleGroup = new THREE.Group();
    this.animationGroup = new THREE.Group();
    this.animationMixer = new THREE.AnimationMixer(this.animationGroup);  
    this.scene.add(this.puzzleGroup);
    this.scene.add(this.animationGroup);
    this.createUiPieces();
  }

  createUiPieces() {
    this.cube.pieces.forEach(piece => {
      const uiPiece = this.createUiPiece(piece);
      this.puzzleGroup.add(uiPiece);
    });
  }

  createAnimationClip(move: Move): THREE.AnimationClip {
    const numTurns = move.numTurns;
    const t0 = 0;
    const t1 = numTurns * (this.animationSpeed / 1000);
    const times = [t0, t1];
    const values: number[] = [];
    const startQuaternion = new THREE.Quaternion();
    const endQuaternion = new THREE.Quaternion();
    const rotationMatrix3 = move.rotationMatrix3;
    const rotationMatrix4 = this.makeRotationMatrix4(rotationMatrix3);
    endQuaternion.setFromRotationMatrix(rotationMatrix4);
    startQuaternion.toArray(values, values.length);
    endQuaternion.toArray(values, values.length);
    const duration = -1;
    const tracks = [new THREE.QuaternionKeyframeTrack(".quaternion", times, values)];
    return new THREE.AnimationClip(move.toString(), duration, tracks);
  }

  animateMoves(moves: Move[], nextMoveIndex: number = 0) {
    const move = moves[nextMoveIndex];
    if (!move) {
      setTimeout(this.scramble.bind(this), AFTER_DELAY);
      return;
    }
    const pieces = this.cube.getPieces(move.coordsList);
    const uiPieces = pieces.map(this.findUiPiece.bind(this));
    this.movePiecesBetweenGroups(uiPieces, this.puzzleGroup, this.animationGroup);
    const onFinished = () => {
      this.animationMixer.removeEventListener("finished", onFinished);
      this.movePiecesBetweenGroups(uiPieces, this.animationGroup, this.puzzleGroup);
      this.cube.pieces = move.makeMove(this.cube.pieces);
      const rotationMatrix3 = move.rotationMatrix3;
      const rotationMatrix4 = this.makeRotationMatrix4(rotationMatrix3);
      for (const uiPiece of uiPieces) {
        uiPiece.applyMatrix4(rotationMatrix4);
      }
      this.animateMoves(moves, nextMoveIndex + 1);
    }

    this.animationMixer.addEventListener("finished", onFinished);

    const animationClip = this.createAnimationClip(move);
    const clipAction = this.animationMixer.clipAction(
      animationClip,
      this.animationGroup);
    clipAction.setLoop(THREE.LoopOnce, 1);
    clipAction.play();
  }

  movePiecesBetweenGroups(uiPieces: THREE.Object3D[], fromGroup: THREE.Group, toGroup: THREE.Group) {
    if (uiPieces.length) {
      fromGroup.remove(...uiPieces);
      toGroup.add(...uiPieces);
    }
  }
}
