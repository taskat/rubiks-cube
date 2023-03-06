import * as THREE from "three"
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls"
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader"
import * as M from "./logic/mymoves"
import * as U from "./logic/utils"
import Cube from "./logic/cube"

const COLOR_TABLE = {
  "U": new THREE.Color("blue"),
  "D": new THREE.Color("green"),
  "L": new THREE.Color("red"),
  "R": new THREE.Color("darkorange"),
  "F": new THREE.Color("yellow"),
  "B": new THREE.Color("ghostwhite"),
  "-": new THREE.Color(0x282828)
}

const PIECE_MATERIAL = new THREE.MeshPhysicalMaterial({
  vertexColors: true,
  metalness: .5,
  roughness: .1,
  clearcoat: .1,
  reflectivity: .5
})

const BEFORE_DELAY = 2000;
const AFTER_DELAY = 2000;

export default class Simulator{
  constructor() {
    this.animationSpeed = 750;
    this.model = null;
    this.modelName = "/cube-bevelled.glb";
    this.cube = null;
    this.cubeSize = 3;
    this.renderer = new THREE.WebGLRenderer({ antialias: true });
    this.camera = new THREE.PerspectiveCamera(34, 1, 1, 100);
    this.scene = new THREE.Scene();
    this.puzzleGroup = new THREE.Group();
    this.animationGroup = new THREE.Group();
    this.clock = new THREE.Clock();
    this.animationMixer = new THREE.AnimationMixer();
  }

  makeRotationMatrix4(rotationMatrix3) {
    let values = new Array(16).fill(0);
    for (let i = 0; i < 3; i++) {
      for (let j = 0; j < 3; j++) {
        values[i * 4 + j] = rotationMatrix3.get([i, j]);
      }
    }
    values[15] = 1;
    return new THREE.Matrix4().fromArray(values);
  }

  initRenderer(container) {
    this.renderer.setPixelRatio(window.devicePixelRatio);
    this.renderer.setSize(container.offsetWidth, container.offsetHeight);
    container.appendChild(this.renderer.domElement);
  }

  addResizeListener(w, h) {
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

  initCamera(w, h) {
    this.camera.aspect = w / h;
    this.camera.position.set(3, 3, 12);
    this.camera.lookAt(new THREE.Vector3(0, 0, 0));
    this.camera.updateProjectionMatrix();
  }

  createLights() {
    let lights = [];
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
    this.controls = new OrbitControls(this.camera, this.renderer.domElement);
    this.controls.minDistance = 5.0;
    this.controls.maxDistance = 40.0;
    this.controls.enableDamping = true;
    this.controls.dampingFactor = 0.9;
    this.controls.autoRotate = false;
    this.controls.autoRotateSpeed = 1.0;
  }

  async init() {
    await this.load3DModel();
    const container = document.getElementById("visualisation-container");
    const w = container.offsetWidth;
    const h = container.offsetHeight;
    this.initRenderer(container);
    this.addResizeListener(w, h);
    this.createScene();
    this.initCamera(w, h);
    this.initControls();
    this.cube = new Cube(this.cubeSize);
    this.createUiPieces();
    this.animate();
    this.scramble();
  }

  async load3DModel() {
    let p =  new Promise((resolve, reject) => {
      const loader = new GLTFLoader();
      loader.load(this.modelName,
        gltf => {
          const bufferGeometry = gltf.scene.children[0].geometry;
          resolve(bufferGeometry.toNonIndexed());
        },
        undefined,
        reject)
    });
    let model = await p;
    this.model = model;
  }

  setGeometryVertexColors(piece) {
    const pieceGeoemtry = this.model.clone();
    const normalAttribute = pieceGeoemtry.getAttribute("normal");

    const colors = [];

    for (let normalIndex = 0; normalIndex < normalAttribute.count; normalIndex += 3) {

      let arrayIndex = normalIndex * normalAttribute.itemSize;
      const normalX = normalAttribute.array[arrayIndex++];
      const normalY = normalAttribute.array[arrayIndex++];
      const normalZ = normalAttribute.array[arrayIndex++];

      const color = this.lookupColorForFaceNormal(piece, normalX, normalY, normalZ);

      colors.push(color.r, color.g, color.b);
      colors.push(color.r, color.g, color.b);
      colors.push(color.r, color.g, color.b);
    }

    pieceGeoemtry.setAttribute("color", new THREE.Float32BufferAttribute(colors, 3));

    return pieceGeoemtry;
  }

  lookupColorForFaceNormal(piece, normalX, normalY, normalZ) {
    if (U.closeTo(normalY, 1)) return COLOR_TABLE[piece.faces.up];
    if (U.closeTo(normalY, -1)) return COLOR_TABLE[piece.faces.down];
    if (U.closeTo(normalX, -1)) return COLOR_TABLE[piece.faces.left];
    if (U.closeTo(normalX, 1)) return COLOR_TABLE[piece.faces.right];
    if (U.closeTo(normalZ, 1)) return COLOR_TABLE[piece.faces.front];
    if (U.closeTo(normalZ, -1)) return COLOR_TABLE[piece.faces.back];
    return COLOR_TABLE["-"];
  }

  animate() {
    window.requestAnimationFrame(this.animate.bind(this));
    this.controls.update();
    const delta = this.clock.getDelta() * this.animationMixer.timeScale;
    this.animationMixer.update(delta);
    this.renderer.render(this.scene, this.camera);
  }

  scramble() {


    // const randomMoves = U.range(NUM_RANDOM_MOVES).map(() => L.getRandomMove(this.cubeSize));
    // L.removeRedundantMoves(randomMoves);
    // console.log(`random moves: ${randomMoves.map(move => move.id).join(" ")}`);
    // globals.cube = L.makeMoves(randomMoves, L.getSolvedCube(this.cubeSize));
    // this.resetUiPieces(globals.cube);
    // setTimeout(showSolutionByCheating, BEFORE_DELAY, randomMoves);
    const moves = M.getMoves(this.cubeSize);
    console.log(`moves: ${moves.map(move => move.id).join(" ")}`);
    this.cube.makeMoves(M.reverseMoves(moves));
    this.resetUiPieces(this.cube);
    setTimeout(this.showSolution.bind(this), BEFORE_DELAY, moves);
  }

  resetUiPieces(cube) {
    cube.pieces.forEach(piece => {
      const uiPiece = this.findUiPiece(piece);
      this.resetUiPiece(uiPiece, piece);
    });
  }

  findUiPiece(piece) {
    return this.puzzleGroup.children.find(child => child.userData === piece.id);
  }

  
  resetUiPiece(uiPiece, piece) {
    const isEvenSizedCube = this.cubeSize % 2 === 0;
    const adjustValue = v => isEvenSizedCube ? v < 0 ? v + 0.5 : v - 0.5 : v;
    uiPiece.position.x = adjustValue(piece.coord.x);
    uiPiece.position.y = adjustValue(piece.coord.y);
    uiPiece.position.z = adjustValue(piece.coord.z);
    uiPiece.setRotationFromMatrix(this.makeRotationMatrix4(piece.accTransform3));
  }

  createUiPiece(piece) {
    const pieceGeometryWithColors = this.setGeometryVertexColors(piece);
    const uiPiece = new THREE.Mesh(pieceGeometryWithColors, PIECE_MATERIAL);
    uiPiece.scale.set(0.5, 0.5, 0.5);
    uiPiece.userData = piece.id;
    this.resetUiPiece(uiPiece, piece);
    return uiPiece;
  }

  recreateUiPieces() {
    this.cube = new Cube(this.cubeSize);
    this.createUiPieces();
  }

  createUiPieces() {
    this.cube.pieces.forEach(piece => {
      const uiPiece = this.createUiPiece(piece);
      this.puzzleGroup.add(uiPiece);
    })
  }

  createAnimationClip(move) {
    const numTurns = move.numTurns;
    const t0 = 0;
    const t1 = numTurns * (this.animationSpeed / 1000);
    const times = [t0, t1];
    const values = [];
    const startQuaternion = new THREE.Quaternion();
    const endQuaternion = new THREE.Quaternion();
    const rotationMatrix3 = move.rotationMatrix3;
    const rotationMatrix4 = this.makeRotationMatrix4(rotationMatrix3);
    endQuaternion.setFromRotationMatrix(rotationMatrix4);
    startQuaternion.toArray(values, values.length);
    endQuaternion.toArray(values, values.length);
    const duration = -1;
    const tracks = [new THREE.QuaternionKeyframeTrack(".quaternion", times, values)];
    return new THREE.AnimationClip(move.id, duration, tracks);
  }


  showSolution(moves) {
    console.log(`solution moves: ${moves.map(move => move.id).join(" ")}`);
    this.animateMoves(moves);
  }

  animateMoves(moves, nextMoveIndex = 0) {

    const move = moves[nextMoveIndex];

    if (!move) {
      return setTimeout(this.scramble.bind(this), AFTER_DELAY);
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
    clipAction.setLoop(THREE.LoopOnce);
    clipAction.play();
  }

  movePiecesBetweenGroups(uiPieces, fromGroup, toGroup) {
    if (uiPieces.length) {
      fromGroup.remove(...uiPieces);
      toGroup.add(...uiPieces);
    }
  }

}
