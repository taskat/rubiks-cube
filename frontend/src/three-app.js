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

const BEFORE_DELAY = 2000;
const AFTER_DELAY = 2000;

export default class Simulator{
  constructor() {
    this.animationSpeed = 750;
    this.model = null;
    this.modelName = "/cube-bevelled.glb";
    this.cube = null;
    this.renderer = new THREE.WebGLRenderer({ antialias: true });
    this.camera = null;
    this.scene = new THREE.Scene();
    this.puzzleGroup = new THREE.Group();
    this.animationGroup = new THREE.Group();
    this.clock = new THREE.Clock();
    this.animationMixer = new THREE.AnimationMixer();
    this.cubeSize = 3;
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
    const isEvenSizedCube = this.cubeSize % 2 === 0;
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
    this.cube = L.getSolvedCube(this.cubeSize);
    this.createUiPieces();
  }

  createUiPieces() {
    this.cube.forEach(piece => {
      const uiPiece = this.createUiPiece(piece);
      this.puzzleGroup.add(uiPiece);
    })
  }

  init() {

    const container = document.getElementById("visualisation-container");
    const w = container.offsetWidth;
    const h = container.offsetHeight;
    this.renderer.setPixelRatio(window.devicePixelRatio);
    this.renderer.setSize(w, h);
    container.appendChild(this.renderer.domElement);

    window.addEventListener("resize", () => {
      this.renderer.setSize(container.offsetWidth, container.offsetHeight);
      this.camera.aspect = container.offsetWidth / container.offsetHeight;
      this.camera.updateProjectionMatrix();
    });

    this.scene.background = new THREE.Color("lightgreen");
    this.camera = new THREE.PerspectiveCamera(34, w / h, 1, 100);
    this.camera.position.set(3, 3, 12);
    this.camera.lookAt(new THREE.Vector3(0, 0, 0));
    this.scene.add(this.camera);

    const LIGHT_COLOR = 0xffffff;
    const LIGHT_INTENSITY = 2;
    const LIGHT_DISTANCE = 4;

    const light1 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light1.position.set(0, 0, LIGHT_DISTANCE);
    this.scene.add(light1);

    const light2 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light2.position.set(0, 0, -LIGHT_DISTANCE);
    this.scene.add(light2);

    const light3 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light3.position.set(0, LIGHT_DISTANCE, 0);
    this.scene.add(light3);

    const light4 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light4.position.set(0, -LIGHT_DISTANCE, 0);
    this.scene.add(light4);

    const light5 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light5.position.set(LIGHT_DISTANCE, 0, 0);
    this.scene.add(light5);

    const light6 = new THREE.DirectionalLight(LIGHT_COLOR, LIGHT_INTENSITY);
    light6.position.set(-LIGHT_DISTANCE, 0, 0);
    this.scene.add(light6)

    
    this.scene.add(this.puzzleGroup);

    
    this.scene.add(this.animationGroup);

    this.controls = new OrbitControls(this.camera, this.renderer.domElement);
    this.controls.minDistance = 5.0;
    this.controls.maxDistance = 40.0;
    this.controls.enableDamping = true;
    this.controls.dampingFactor = 0.9;
    this.controls.autoRotate = false;
    this.controls.autoRotateSpeed = 1.0;

    
    

    this.cube = L.getSolvedCube(this.cubeSize);
    
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
    this.cube = L.makeMoves(M.reverseMoves(moves), L.getSolvedCube(this.cubeSize));
    this.resetUiPieces(this.cube);
    setTimeout(this.showSolution.bind(this), BEFORE_DELAY, moves);
  }

  resetUiPieces(cube) {
    cube.forEach(piece => {
      const uiPiece = this.findUiPiece(piece);
      this.resetUiPiece(uiPiece, piece);
    });
  }

  findUiPiece(piece) {
    return this.puzzleGroup.children.find(child => child.userData === piece.id);
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

    const pieces = L.getPieces(this.cube, move.coordsList);
    const uiPieces = pieces.map(this.findUiPiece.bind(this));
    this.movePiecesBetweenGroups(uiPieces, this.puzzleGroup, this.animationGroup);

    const onFinished = () => {
      this.animationMixer.removeEventListener("finished", onFinished);
      this.movePiecesBetweenGroups(uiPieces, this.animationGroup, this.puzzleGroup);
      this.cube = move.makeMove(this.cube);
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
