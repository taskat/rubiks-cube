import * as THREE from "three"
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls"
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader"
import * as L from "./logic"
import * as M from "./logic/mymoves"
import * as U from "./logic/utils"

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

const globals = {
  pieceGeometry: undefined,
  cube: undefined,
  renderer: undefined,
  camera: undefined,
  scene: undefined,
  puzzleGroup: undefined,
  animationGroup: undefined,
  axesHelper: undefined,
  controls: undefined,
  clock: undefined,
  animationMixer: undefined,
  cubeSize: 3,
  cubeSizeChanged: true,
  axesEnabled: false
}
const BEFORE_DELAY = 2000;
const AFTER_DELAY = 2000;

export default class Simulator{
  constructor() {
    this.animationSpeed = 750;
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

  resetUiPiece(uiPiece, piece) {
    const isEvenSizedCube = globals.cubeSize % 2 === 0;
    const adjustValue = v => isEvenSizedCube ? v < 0 ? v + 0.5 : v - 0.5 : v;
    uiPiece.position.x = adjustValue(piece.x);
    uiPiece.position.y = adjustValue(piece.y);
    uiPiece.position.z = adjustValue(piece.z);
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
    globals.cube = L.getSolvedCube(globals.cubeSize);
    this.createUiPieces();
  }

  createUiPieces() {
    globals.cube.forEach(piece => {
      const uiPiece = this.createUiPiece(piece);
      globals.puzzleGroup.add(uiPiece);
    })
  }

  async init() {

    const container = document.getElementById("visualisation-container");
    const w = container.offsetWidth;
    const h = container.offsetHeight;
    globals.renderer = new THREE.WebGLRenderer({ antialias: true });
    globals.renderer.setPixelRatio(window.devicePixelRatio);
    globals.renderer.setSize(w, h);
    container.appendChild(globals.renderer.domElement);

    window.addEventListener("resize", () => {
      globals.renderer.setSize(container.offsetWidth, container.offsetHeight);
      globals.camera.aspect = container.offsetWidth / container.offsetHeight;
      globals.camera.updateProjectionMatrix();
    });

    globals.scene = new THREE.Scene();
    globals.scene.background = new THREE.Color("lightgreen");
    globals.camera = new THREE.PerspectiveCamera(34, w / h, 1, 100);
    globals.camera.position.set(3, 3, 12);
    globals.camera.lookAt(new THREE.Vector3(0, 0, 0));
    globals.scene.add(globals.camera);

    const LIGHT_COLOR = 0xffffff;
    const LIGHT_INTENSITY = 2;
    const LIGHT_DISTANCE = 4;

    const light1 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light1.position.set(0, 0, LIGHT_DISTANCE);
    globals.scene.add(light1);

    const light2 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light2.position.set(0, 0, -LIGHT_DISTANCE);
    globals.scene.add(light2);

    const light3 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light3.position.set(0, LIGHT_DISTANCE, 0);
    globals.scene.add(light3);

    const light4 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light4.position.set(0, -LIGHT_DISTANCE, 0);
    globals.scene.add(light4);

    const light5 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light5.position.set(LIGHT_DISTANCE, 0, 0);
    globals.scene.add(light5);

    const light6 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light6.position.set(-LIGHT_DISTANCE, 0, 0);
    globals.scene.add(light6)

    globals.puzzleGroup = new THREE.Group();
    globals.scene.add(globals.puzzleGroup);

    globals.animationGroup = new THREE.Group();
    globals.scene.add(globals.animationGroup);

    globals.controls = new OrbitControls(globals.camera, globals.renderer.domElement);
    globals.controls.minDistance = 5.0;
    globals.controls.maxDistance = 40.0;
    globals.controls.enableDamping = true;
    globals.controls.dampingFactor = 0.9;
    globals.controls.autoRotate = false;
    globals.controls.autoRotateSpeed = 1.0;

    globals.clock = new THREE.Clock();
    globals.animationMixer = new THREE.AnimationMixer();

    globals.cube = L.getSolvedCube(globals.cubeSize);
    globals.pieceGeometry = await this.load3DModel("/cube-bevelled.glb");
    this.createUiPieces();

    this.animate();
    this.scramble();
  }

  load3DModel(url) {
    return new Promise((resolve, reject) => {
      const loader = new GLTFLoader()
      loader.load(
        url,
        gltf => {
          const bufferGeometry = gltf.scene.children[0].geometry
          resolve(bufferGeometry.toNonIndexed())
        },
        undefined,
        reject)
    });
  }

  setGeometryVertexColors(piece) {
    const pieceGeoemtry = globals.pieceGeometry.clone();
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
    globals.controls.update();
    const delta = globals.clock.getDelta() * globals.animationMixer.timeScale;
    globals.animationMixer.update(delta);
    globals.renderer.render(globals.scene, globals.camera);
  }

  scramble() {

    if (globals.cubeSizeChanged) {
      globals.cubeSizeChanged = false;
      globals.puzzleGroup.clear();
      globals.animationGroup.clear();
      globals.controls.reset();
      const cameraX = globals.cubeSize + 1;
      const cameraY = globals.cubeSize + 1;
      const cameraZ = globals.cubeSize * 4;
      globals.camera.position.set(cameraX, cameraY, cameraZ);
      globals.camera.lookAt(new THREE.Vector3(0, 0, 0));
      this.recreateUiPieces();
    }

    // const randomMoves = U.range(NUM_RANDOM_MOVES).map(() => L.getRandomMove(globals.cubeSize));
    // L.removeRedundantMoves(randomMoves);
    // console.log(`random moves: ${randomMoves.map(move => move.id).join(" ")}`);
    // globals.cube = L.makeMoves(randomMoves, L.getSolvedCube(globals.cubeSize));
    // this.resetUiPieces(globals.cube);
    // setTimeout(showSolutionByCheating, BEFORE_DELAY, randomMoves);
    const moves = M.getMoves(globals.cubeSize);
    console.log(`moves: ${moves.map(move => move.id).join(" ")}`);
    globals.cube = L.makeMoves(M.reverseMoves(moves), L.getSolvedCube(globals.cubeSize));
    this.resetUiPieces(globals.cube);
    setTimeout(this.showSolution.bind(this), BEFORE_DELAY, moves);
  }

  resetUiPieces(cube) {
    cube.forEach(piece => {
      const uiPiece = this.findUiPiece(piece);
      this.resetUiPiece(uiPiece, piece);
    });
  }

  findUiPiece(piece) {
    return globals.puzzleGroup.children.find(child => child.userData === piece.id);
  }

  showSolution(moves) {
    console.log(`solution moves: ${moves.map(move => move.id).join(" ")}`);
    this.animateMoves(moves);
  }

  animateMoves(moves, nextMoveIndex = 0) {

    if (globals.cubeSizeChanged) {
      return setTimeout(scramble, 0);
    }

    const move = moves[nextMoveIndex];

    if (!move) {
      return setTimeout(this.scramble.bind(this), AFTER_DELAY);
    }

    const pieces = L.getPieces(globals.cube, move.coordsList);
    const uiPieces = pieces.map(this.findUiPiece.bind(this));
    this.movePiecesBetweenGroups(uiPieces, globals.puzzleGroup, globals.animationGroup);

    const onFinished = () => {
      globals.animationMixer.removeEventListener("finished", onFinished);
      this.movePiecesBetweenGroups(uiPieces, globals.animationGroup, globals.puzzleGroup);
      globals.cube = move.makeMove(globals.cube);
      const rotationMatrix3 = move.rotationMatrix3;
      const rotationMatrix4 = this.makeRotationMatrix4(rotationMatrix3);
      for (const uiPiece of uiPieces) {
        uiPiece.applyMatrix4(rotationMatrix4);
      }
      this.animateMoves(moves, nextMoveIndex + 1);
    }

    globals.animationMixer.addEventListener("finished", onFinished);

    const animationClip = this.createAnimationClip(move);
    const clipAction = globals.animationMixer.clipAction(
      animationClip,
      globals.animationGroup);
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
